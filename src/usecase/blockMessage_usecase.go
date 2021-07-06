package usecase

import (
	"context"
	"message-service/domain"
	"message-service/repository"
)

type blockMessageUsecase struct {
	BlockMessageRepository repository.BlockMessageRepository
}

type BlockMessageUsecase interface {
	IsBlocked(ctx context.Context, blockedBy, blockedFor string) (bool, error)
	BlockMessage(ctx context.Context, blockedBy, blockedFor string) (*domain.BlockMessage, error)
}

func NewBlockMessageUsecase(blockMessageRepository repository.BlockMessageRepository) BlockMessageUsecase {
	return &blockMessageUsecase{
		BlockMessageRepository: blockMessageRepository,
	}
}

func (b *blockMessageUsecase) IsBlocked(ctx context.Context, blockedBy, blockedFor string) (bool, error) {
	return b.BlockMessageRepository.IsBlocked(ctx, blockedBy, blockedFor)
}

func (b *blockMessageUsecase) BlockMessage(ctx context.Context, blockedBy, blockedFor string) (*domain.BlockMessage, error) {
	return b.BlockMessageRepository.BlockMessage(ctx, blockedBy, blockedFor)
}



