package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"message-service/domain"
	"message-service/infrastructure/chat"
	"message-service/infrastructure/mongo"
	"message-service/infrastructure/seeder"
	"message-service/repository"
	"message-service/usecase"
)

func main() {


	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()
	seeder.SeedData(db, mongoCli, *ctx)



	messageRepository := repository.NewMessageRepository(mongoCli)
	messageUsecase := usecase.NewMessageUsecase(messageRepository)
	blockMessageRepository := repository.NewBlockMessageRepository(mongoCli)
	blockMessageUsecase := usecase.NewBlockMessageUsecase(blockMessageRepository)
	messageRequestRepository := repository.NewMessageRequestRepository(mongoCli)
	messageRequestUsecase := usecase.MessageRequestUsecase(messageRequestRepository)


	hub := chat.NewHub(messageUsecase, blockMessageUsecase, messageRequestUsecase)
	chatClient := chat.NewClient(hub)
	go hub.Run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/message/:receiver/:sender", func(c *gin.Context) {

		receiverId := c.Param("receiver")
		senderId := c.Param("sender")


		messages, err := messageUsecase.GetMessages(c, receiverId, senderId)

		if err != nil {
			c.JSON(400, gin.H{"message" : "Error getting messages"})
			return
		}

		c.JSON(200, messages)
	})

	router.GET("/messages/:userId", func(c *gin.Context) {
		userId := c.Param("userId")


		messages, err := messageUsecase.GetFirstMessages(c, userId)

		if err != nil {
			c.JSON(400, gin.H{"message" : "Error getting messages"})
			return
		}

		c.JSON(200, messages)

	})

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.POST("/messageRequest", func(c *gin.Context) {
		var message domain.Message

		decoder := json.NewDecoder(c.Request.Body)
		if err := decoder.Decode(&message); err != nil {
			c.JSON(400, gin.H{"message" : "error decoding body"})
			return
		}

		if _, err := messageRequestUsecase.Create(c, message); err != nil {
			c.JSON(400, gin.H{"message" : "error creating request"})
			return
		}

	})
	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		chatClient.ServeWs(c.Writer, c.Request, roomId)
	})



	router.Run("127.0.0.1:8080")

}
