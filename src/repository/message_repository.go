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
	GetUsers(ctx context.Context, userId string) ([]string, error)
	UpdateSeenStatus(ctx context.Context, messageId string, value bool) error
	IsAllowedToSee(ctx context.Context, messageId string) bool
	Delete(ctx context.Context, messageId string) error
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

	filter2 := bson.D{{"$or", []bson.D{
		bson.D{{"message_to._id", sender}, {"message_from._id",receiver}, {"seen" , false}},
		bson.D{{"message_to._id",receiver}, {"message_from._id", sender}, {"seen" , false}},
	}}}


	filter, err := m.collection.Find(ctx, filter2, findOptions)
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

func (m *messageRepository) GetUsers(ctx context.Context, userId string) ([]string, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	findOptions := options.Find()

	findOptions.SetLimit(1)

	var userIds []string

	results, err := m.collection.Distinct(ctx, "message_from._id", bson.M{"message_to._id" : userId})

	if err != nil {
		return nil, err
	}

	for _, res := range results {
		userId := res.(string)
		userIds = append(userIds, userId)
	}



	return userIds, nil
}

func (m *messageRepository) UpdateSeenStatus(ctx context.Context, messageId string, value bool) error {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()


	messageToUpdate := bson.M{"_id" : messageId}

	updatedMessage := bson.M{"$set" : bson.M{
		"seen" : value,
		"image_path" : "",
	}}

	_, err := m.collection.UpdateOne(ctx, messageToUpdate, updatedMessage)

	if err != nil {
		return err
	}

	return nil
}

func (m *messageRepository) IsAllowedToSee(ctx context.Context, messageId string) bool {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	filter := bson.M{"_id" : messageId}

	var result domain.Message
	if err := m.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return false
	}

	if result.Seen {
		return false
	}
	updatedValue := bson.M{"$set" : bson.M{
		"seen" : true,
	}}

	if _, err := m.collection.UpdateOne(ctx, filter, updatedValue); err != nil {
		return false
	}

	return true
}

func (m *messageRepository) Delete(ctx context.Context, messageId string) error {

	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	filter := bson.M{"_id" : messageId}

	if _, err := m.collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

