package domain

import (
	"message-service/domain/enum"
	"time"
)

type Message struct {
	Id		uint64`json:"id"`
	Timestamp		time.Time `json:"timestamp"`
	Content		string `json:"content"`
	Path	string `json:"redirect_path"`
	Type 	enum.MessageType `json:"type"`
}
