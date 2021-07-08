package usecase

import (
	"context"
	"message-service/domain"
	"message-service/repository"
	"time"
)

type messageRequestUsecase struct {
	MessageRequestRepository repository.MessageRequestRepository
	MessageRepository repository.MessageRepository
}


type MessageRequestUsecase interface {
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
	IsCreated(ctx context.Context, messaege domain.Message) (bool, error)
	GetByUserId(ctx context.Context, userId string) ([]*domain.Message, error)
	AcceptRequest(ctx context.Context, messageRequestId string) error
	RejectRequest(ctx context.Context, messageRequestId string) error
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

func (m *messageRequestUsecase) AcceptRequest(ctx context.Context, messageRequestId string) error {

	deletedRequest, err := m.MessageRequestRepository.Delete(ctx, messageRequestId)
	if err != nil {
		return err
	}

	deletedRequest.Timestamp = time.Now()
	if _, err:= m.MessageRepository.Create(ctx, *deletedRequest); err != nil {
		return err
	}

	return nil
}

func (m *messageRequestUsecase) RejectRequest(ctx context.Context, messageRequestId string) error {

	if _, err := m.MessageRequestRepository.Delete(ctx, messageRequestId); err != nil {
		return err
	}

	return nil
}

