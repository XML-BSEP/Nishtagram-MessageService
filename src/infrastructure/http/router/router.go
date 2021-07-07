package router

import (
	"github.com/gin-gonic/gin"
	"message-service/infrastructure/http/middleware"
	"message-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.LoadHTMLFiles("index.html")

	router.GET("/message/:receiver/:sender", handler.GetMessages)
	router.GET("/messageRequest/:userId", handler.GetMessageRequests)
	router.POST("/blockMessage", handler.Block)
	//router.GET("/ws/:roomId", handler.Serve)
	router.GET("/room/:roomId", handler.GetRoom)
	router.GET("/users/:userId", handler.GetUsers)
	router.GET("/blocked/:blockedBy/:blockedFor", handler.IsBlocked)


	return router
}
