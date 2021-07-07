package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dto2 "message-service/infrastructure/dto"
	"message-service/usecase"
)

type blockMessageHandler struct {
	BlockMessageUsecase usecase.BlockMessageUsecase
}

type BlockMessageHandler interface {
	Block(c *gin.Context)
}

func NewBlockMessageUsecase(blockMessageUsecase usecase.BlockMessageUsecase) BlockMessageHandler {
	return &blockMessageHandler{BlockMessageUsecase: blockMessageUsecase}
}

func (b *blockMessageHandler) Block(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	var dto dto2.BlockMessageDto

	if err := decoder.Decode(&dto); err != nil {
		c.JSON(500, gin.H{"message" : "Error decoding body"})
		return
	}

	if dto.BlockedFor == "" || dto.BlockedBy == "" {
		c.JSON(400, gin.H{"message" : "Error blocking messages"})
		return
	}
	if _, err := b.BlockMessageUsecase.BlockMessage(c, dto.BlockedBy, dto.BlockedFor); err != nil {
		c.JSON(400, gin.H{"message" : "Error blocking messages"})
		return
	}

	c.JSON(200, gin.H{"message" : "Successfully blocked"})

}