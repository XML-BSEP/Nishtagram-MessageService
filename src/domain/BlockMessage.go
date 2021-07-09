package domain

import "time"

type BlockMessage struct {
	ID         string    `bson:"_id" json:"id"`
	Timestamp  time.Time `bson:"timestamp" json:"timestamp"`
	BlockedBy  Profile   `bson:"blocked_by" json:"blocked_by"`
	BlockedFor Profile   `bson:"blocked_for" json:"blocked_for"`
}
