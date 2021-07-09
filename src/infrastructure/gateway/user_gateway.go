package gateway

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"message-service/infrastructure/dto"
	"os"
)

func GetUserInfo(context context.Context, ids dto.UserIdsDto) ([]dto.UserDto, error) {
	client := resty.New()

	domain := os.Getenv("MESSAGE_DOMAIN")

	if domain == "" {
		domain = "127.0.0.1"
	}

	var responseDto []dto.UserDto

	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(ids).
			EnableTrace().
			SetContext(context).
			Post("https://" + domain + ":8082/getSearchInfo")

		err := json.Unmarshal(resp.Body(), &responseDto)

		if err != nil {
			return nil, err
		}
		return responseDto, nil
	} else {
		resp, _ := client.R().
			SetBody(ids).
			EnableTrace().
			SetContext(context).
			Post("http://" + domain + ":8082/getSearchInfo")

		err := json.Unmarshal(resp.Body(), &responseDto)

		if err != nil {
			return nil, err
		}
		return responseDto, nil
	}


}
