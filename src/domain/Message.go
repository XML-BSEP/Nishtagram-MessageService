package domain

import (
	"message-service/domain/enum"
	"time"
)

type Message struct {
	ID          string           `bson:"_id" json:"id"`
	Timestamp   time.Time        `bson:"timestamp" json:"timestamp"`
	Content     string           `bson:"content" json:"content"`
	Path        string           `bson:"path" json:"redirect_path"`
	Type        enum.MessageType `bson:"type" json:"type"`
	MessageFrom Profile          `bson:"message_from" json:"message_from"`
	MessageTo   Profile          `bson:"message_to" json:"message_to"`
	ImagePath   string `bson:"image_path" json:"image_path"`
	ImageBase64 string `json:"image_base_64"`
	Seen bool `bson:"seen" json:"seen"`
	ResponseType enum.ResponseType `json:"response_type"`
}
