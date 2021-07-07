package handler

import (
	"github.com/gin-gonic/gin"
	"message-service/usecase"
)

type messageHandler struct {
	MessageUsecase usecase.MessageUsecase

}


type MessageHandler interface {
	GetMessages(c *gin.Context)
	GetUsers(c *gin.Context)
}

func NewMessageHandler(messageUsecase usecase.MessageUsecase) MessageHandler {
	return &messageHandler{MessageUsecase: messageUsecase}
}

func (m *messageHandler) GetMessages(c *gin.Context) {
	receiverId := c.Param("receiver")
	senderId := c.Param("sender")

	messages, err := m.MessageUsecase.GetMessages(c, receiverId, senderId)

	if err != nil {
		c.JSON(400, gin.H{"message": "Error getting messages"})
		return
	}

	c.JSON(200, messages)
}

func (m *messageHandler) GetUsers(c *gin.Context) {

	userId := c.Param("userId")

	messages, err := m.MessageUsecase.GetUsers(c, userId)

	if err != nil {
		c.JSON(400, gin.H{"message" : "There is no messages"})
		return
	}

	c.JSON(200, messages)

}



