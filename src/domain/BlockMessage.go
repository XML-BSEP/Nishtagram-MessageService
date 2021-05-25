package domain

import "time"

type BlockMessage struct {
	Id		uint64`json:"id"`
	Timestamp	time.Time `json:"timestamp"`
	BlockMessagesFor 	Profile `json:"block_notification_for"`
	BlockedMessages	Profile `json:"blocked_notification"`

}