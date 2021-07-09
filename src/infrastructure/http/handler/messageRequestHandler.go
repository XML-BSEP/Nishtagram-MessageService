package handler

import (
	"github.com/gin-gonic/gin"
	"message-service/usecase"
)

type messageRequestHandler struct {
	MessageRequestUsecase usecase.MessageRequestUsecase
}

type MessageRequestHandler interface {
	GetMessageRequests(c *gin.Context)
}

func NewMessageRequestHandler(messageRequestUsecase usecase.MessageRequestUsecase) MessageRequestHandler {
	return &messageRequestHandler{MessageRequestUsecase: messageRequestUsecase}
}

func (m *messageRequestHandler) GetMessageRequests(c *gin.Context) {

	userId := c.Param("userId")

	messageRequests, err := m.MessageRequestUsecase.GetByUserId(c, userId)

	if err != nil {
		c.JSON(400, gin.H{"message" : "Error getting message requests"})
		return
	}

	c.JSON(200, messageRequests)
}

