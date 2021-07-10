package gateway

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"message-service/infrastructure/dto"
	"os"
)

func IsFollowing(context context.Context, users dto.FollowDTO) (bool, error) {
	client := resty.New()

	domain := os.Getenv("FOLLOW_DOMAIN")

	if domain == "" {
		domain = "127.0.0.1"
	}

	var response bool

	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(users).
			EnableTrace().
			SetContext(context).
			Post("https://" + domain + ":8089/isUserFollowingUser")

		err := json.Unmarshal(resp.Body(), &response)

		if err != nil {
			return false, err
		}
		return response, nil
	} else {
		resp, _ := client.R().
			SetBody(users).
			EnableTrace().
			SetContext(context).
			Post("http://" + domain + ":8089/isUserFollowingUser")

		err := json.Unmarshal(resp.Body(), &response)

		if err != nil {
			return false, err
		}
		return response, nil
	}


}
