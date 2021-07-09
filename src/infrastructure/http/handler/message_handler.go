package handler

import (
	"github.com/gin-gonic/gin"
	"message-service/infrastructure/dto"
	"message-service/infrastructure/gateway"
	"message-service/usecase"
)

type messageHandler struct {
	MessageUsecase usecase.MessageUsecase

}


type MessageHandler interface {
	GetMessages(c *gin.Context)
	GetUsers(c *gin.Context)
	IsAllowedToSee(c *gin.Context)
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

	idsDto := dto.UserIdsDto{Ids: messages}
	users, err := gateway.GetUserInfo(c, idsDto)

	if err != nil {
		c.JSON(400, gin.H{"message" : "Can not get user infos"})
		return
	}

	c.JSON(200, users)

}

func (m *messageHandler) IsAllowedToSee(c *gin.Context) {

	messageId := c.Param("messageId")

	if isAllowed := m.MessageUsecase.IsAllowedToSee(c, messageId); !isAllowed {
		c.JSON(400, gin.H{"message" : "Already seen"})
		return
	}

	c.JSON(200, gin.H{})
}



