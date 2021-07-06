package enum

type MessageType int

const (
	Text	MessageType = iota
	Media
	Bomb
)
