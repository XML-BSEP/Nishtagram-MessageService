package usecase

import (
	"context"
	"fmt"
	"message-service/domain"
	"message-service/repository"
)

type messageUsecsae struct {
	MessageRepository repository.MessageRepository
}


type MessageUsecase interface {
	GetMessages(ctx context.Context, receiver, sender string) ([]*domain.Message, error)
	GetUsers(ctx context.Context, userId string) ([]string, error)
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
}

func NewMessageUsecase(messageRepository repository.MessageRepository) MessageUsecase {
	return &messageUsecsae{
		MessageRepository: messageRepository,
	}
}

func (m *messageUsecsae) GetMessages(ctx context.Context, receiver, sender string) ([]*domain.Message, error) {
	return m.MessageRepository.GetMessages(ctx, receiver, sender)
}

func (m *messageUsecsae) GetUsers(ctx context.Context, userId string) ([]string, error) {
	return m.MessageRepository.GetUsers(ctx, userId)
}

func (m *messageUsecsae) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	fmt.Println("Usao u create")
	return m.MessageRepository.Create(ctx, message)
}

