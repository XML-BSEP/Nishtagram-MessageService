package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"message-service/domain"
	"time"
)

type blockMessageRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}


type BlockMessageRepository interface {
	IsBlocked(ctx context.Context, blockedBy, blockedFor string) (bool, error)
	BlockMessage(ctx context.Context, blockedBy, blockedFor string) (*domain.BlockMessage, error)
}

func NewBlockMessageRepository(db *mongo.Client) BlockMessageRepository {
	return &blockMessageRepository{
		db: db,
		collection: db.Database("message_db").Collection("blocked_messages"),
	}
}

func (b *blockMessageRepository) IsBlocked(ctx context.Context, blockedBy, blockedFor string) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	if err := b.collection.FindOne(ctx, bson.M{"blocked_by._id" : blockedBy, "blocked_for._id" : blockedFor}).Err(); err != nil {
		return false, err
	}

	return true, nil
}



func (b *blockMessageRepository) BlockMessage(ctx context.Context, blockedBy, blockedFor string) (*domain.BlockMessage, error) {

	if blocked, _ := b.IsBlocked(ctx, blockedBy, blockedFor); blocked {
		return nil, fmt.Errorf("Already blocked")
	}

	blockMessage := domain.BlockMessage{
		ID: uuid.NewString(),
		Timestamp: time.Now(),
		BlockedBy: domain.Profile{ID: blockedBy},
		BlockedFor: domain.Profile{ID: blockedFor},
	}

	if _, err := b.collection.InsertOne(ctx, blockMessage); err != nil {
		return nil, err
	}

	return &blockMessage, nil
}


