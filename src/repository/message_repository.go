package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"message-service/domain"
	"time"
)

type messageRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}


type MessageRepository interface {

	GetMessages(ctx context.Context, receiver, sender string) ([]*domain.Message, error)
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
	GetFirstMessages(ctx context.Context, userId string) ([]*domain.Message, error)
}

func NewMessageRepository(db *mongo.Client) MessageRepository {
	return &messageRepository{
		db: db,
		collection: db.Database("message_db").Collection("messages"),
	}
}

func (m *messageRepository) GetMessages(ctx context.Context, receiver, sender string) ([]*domain.Message, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var messages []*domain.Message
	findOptions := options.Find()

	findOptions.SetSort(map[string]int{"timestamp" : 1})

	filter, err := m.collection.Find(ctx, bson.M{"message_to._id" : receiver, "message_from._id" : sender}, findOptions)

	if err != nil {
		return nil, err
	}

	if err := filter.All(ctx, &messages); err != nil {
		return nil, err
	}



	return messages, nil
}

func (m *messageRepository) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()


	if _, err := m.collection.InsertOne(ctx, message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (m *messageRepository) GetFirstMessages(ctx context.Context, userId string) ([]*domain.Message, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	findOptions := options.Find()

	findOptions.SetLimit(1)

	var messages []*domain.Message

	results, err := m.collection.Find(ctx, bson.M{"message_to._id" : userId}, findOptions)

	if err != nil {
		return nil, err
	}

	if err := results.All(ctx, &messages); err != nil {
		return nil, err
	}


	return messages, nil
}

