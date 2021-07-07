package dto

type BlockMessageDto struct {
	BlockedBy string `json:"blocked_by"`
	BlockedFor string `json:"blocked_for"`
}
