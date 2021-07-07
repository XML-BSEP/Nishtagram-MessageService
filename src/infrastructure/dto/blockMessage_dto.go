package dto

type BlockMessageDto struct {
	BlockedBy string `json:"blocked_by"`
	BlockedFor string `json:"blocked_for"`
}

type BlockDto struct {
	IsBlocked bool `json:"is_blocked"`
}
