package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"message-service/domain"
	"time"
)

type messageRequestRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}

type MessageRequestRepository interface {
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
	IsCreated(ctx context.Context, messaege domain.Message) (bool, error)
}

func NewMessageRequestRepository(db *mongo.Client) MessageRequestRepository {
	return &messageRequestRepository{
		db: db,
		collection: db.Database("message_db").Collection("message_requests"),
	}
}

func (m *messageRequestRepository) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if created, _ := m.IsCreated(ctx, message); created {
		return nil, fmt.Errorf("Already created")
	}

	if _, err := m.collection.InsertOne(ctx, message); err != nil {
		return nil, err
	}

	return &message, nil

}

func (m *messageRequestRepository) IsCreated(ctx context.Context, message domain.Message) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := m.collection.FindOne(ctx, bson.M{"message_from._id" : message.MessageFrom.ID, "blocked_for._id" : message.MessageTo.ID}).Err(); err != nil {
		return false, err
	}

	return true, nil

}




