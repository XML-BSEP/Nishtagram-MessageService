package handler

import (
	"github.com/gin-gonic/gin"
	"message-service/infrastructure/chat"
)

type chatHandler struct {
	chatClient chat.Client
}


type ChatHandler interface {
	GetRoom(c *gin.Context)
}

func NewChatHandler(c chat.Client) ChatHandler {
	return &chatHandler{chatClient: c}
}


func (ch *chatHandler) GetRoom(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}


