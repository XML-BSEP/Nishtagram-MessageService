package enum

type MessageType int

const (
	Text	MessageType = iota + 1
	Media
	Bomb
)
