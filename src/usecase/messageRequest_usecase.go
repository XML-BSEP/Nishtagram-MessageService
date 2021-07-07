package usecase

import (
	"context"
	"message-service/domain"
	"message-service/repository"
)

type messageRequestUsecase struct {
	MessageRequestRepository repository.MessageRequestRepository
}

type MessageRequestUsecase interface {
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
	IsCreated(ctx context.Context, messaege domain.Message) (bool, error)
	GetByUserId(ctx context.Context, userId string) ([]*domain.Message, error)
}

func NewMessageRequestUsecase(messageRequestRepository repository.MessageRequestRepository) MessageRequestUsecase{
	return &messageRequestUsecase{
		MessageRequestRepository: messageRequestRepository,
	}
}

func (m *messageRequestUsecase) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	return m.MessageRequestRepository.Create(ctx, message)
}

func (m *messageRequestUsecase) IsCreated(ctx context.Context, messaege domain.Message) (bool, error) {
	return m.IsCreated(ctx, messaege)
}

func (m *messageRequestUsecase) GetByUserId(ctx context.Context, userId string) ([]*domain.Message, error) {
	return m.MessageRequestRepository.GetByUserId(ctx, userId)
}

