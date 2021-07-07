package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"message-service/infrastructure/chat"
	"message-service/infrastructure/http/handler"
	"message-service/repository"
	"message-service/usecase"
)

type interactor struct {
	db *mongo.Client

}

type Interactor interface {
	NewMessageRepository() repository.MessageRepository
	NewBlockMessageRepository() repository.BlockMessageRepository
	NewMessageRequestRepository() repository.MessageRequestRepository

	NewMessageUsecase() usecase.MessageUsecase
	NewBlockMessageUsecase() usecase.BlockMessageUsecase
	NewMessageRequestUsecase() usecase.MessageRequestUsecase

	NewMessageHandler() handler.MessageHandler
	NewMessageRequestHandler() handler.MessageRequestHandler
	NewBlockMessageHandler() handler.BlockMessageHandler
	NewChatHandler() handler.ChatHandler

	NewChatHub() chat.Hub
	NewChatClient() chat.Client

	NewAppHandler() AppHandler
}

func NewInteractor(db *mongo.Client) Interactor {
	return &interactor{db: db}
}

func (i *interactor) NewMessageRepository() repository.MessageRepository {
	return repository.NewMessageRepository(i.db)
}

func (i *interactor) NewBlockMessageRepository() repository.BlockMessageRepository {
	return repository.NewBlockMessageRepository(i.db)
}

func (i *interactor) NewMessageRequestRepository() repository.MessageRequestRepository {
	return repository.NewMessageRequestRepository(i.db)
}

func (i *interactor) NewMessageUsecase() usecase.MessageUsecase {
	return usecase.NewMessageUsecase(i.NewMessageRepository())
}

func (i *interactor) NewBlockMessageUsecase() usecase.BlockMessageUsecase {
	return usecase.NewBlockMessageUsecase(i.NewBlockMessageRepository())
}

func (i *interactor) NewMessageRequestUsecase() usecase.MessageRequestUsecase {
	return usecase.NewMessageRequestUsecase(i.NewMessageRequestRepository())
}

func (i *interactor) NewMessageHandler() handler.MessageHandler {
	return handler.NewMessageHandler(i.NewMessageUsecase())
}

func (i *interactor) NewMessageRequestHandler() handler.MessageRequestHandler {
	return handler.NewMessageRequestHandler(i.NewMessageRequestUsecase())
}

func (i *interactor) NewBlockMessageHandler() handler.BlockMessageHandler {
	return handler.NewBlockMessageUsecase(i.NewBlockMessageUsecase())
}

func (i *interactor) NewChatHandler() handler.ChatHandler {
	return handler.NewChatHandler(i.NewChatClient())
}


func (i *interactor) NewChatHub() chat.Hub {
	return chat.NewHub(i.NewMessageUsecase(), i.NewBlockMessageUsecase(), i.NewMessageRequestUsecase())
}

func (i *interactor) NewChatClient() chat.Client {
	return chat.NewClient(i.NewChatHub())
}




type appHandler struct {
	handler.MessageHandler
	handler.MessageRequestHandler
	handler.BlockMessageHandler
	handler.ChatHandler
}

type AppHandler interface {
	handler.MessageHandler
	handler.MessageRequestHandler
	handler.BlockMessageHandler
	handler.ChatHandler
}

func (i *interactor) NewAppHandler() AppHandler {
	appHandler := &appHandler{}
	appHandler.MessageHandler = i.NewMessageHandler()
	appHandler.MessageRequestHandler = i.NewMessageRequestHandler()
	appHandler.BlockMessageHandler = i.NewBlockMessageHandler()
	appHandler.ChatHandler = i.NewChatHandler()

	return appHandler
}

