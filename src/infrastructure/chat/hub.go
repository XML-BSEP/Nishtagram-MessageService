package chat

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"message-service/domain"
	"message-service/usecase"
	"time"
)

type Hub struct {
	rooms map[string]map[*connection]bool
	connections map[string]bool
	broadcast chan message
	register chan Subscription
	unregister chan Subscription
	MessageUsecase usecase.MessageUsecase
	BlockMessageUsecase usecase.BlockMessageUsecase
	MessageRequestUsecase usecase.MessageRequestUsecase
}

func NewHub(messageUsecase usecase.MessageUsecase, blockMessageUsecase usecase.BlockMessageUsecase, messageRequestUsecase usecase.MessageRequestUsecase) Hub {
	return Hub{
		broadcast:  make(chan message),
		register:   make(chan Subscription),
		unregister: make(chan Subscription),
		rooms:      make(map[string]map[*connection]bool),
		connections: make(map[string]bool),
		MessageUsecase: messageUsecase,
		BlockMessageUsecase: blockMessageUsecase,
		MessageRequestUsecase: messageRequestUsecase,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <- h.register:


			connections := h.rooms[s.room]
			if connections == nil {
				connections := make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true

		case s := <- h.unregister:
			connections := h.rooms[s.room]
			if connections == nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <- h.broadcast:
			connections := h.rooms[m.room]
			var message domain.Message
			if err := json.Unmarshal(m.data, &message); err != nil {
				break
			}

			if blocked, _  := h.BlockMessageUsecase.IsBlocked(context.Background(), message.MessageTo.ID, message.MessageFrom.ID); blocked {
				break
			}
			message.ID = uuid.NewString()
			message.Timestamp = time.Now()
			h.MessageUsecase.Create(context.Background(), message)
			for c := range connections {
				select {
				case c.send <- m.data:
					/*message := domain.Message{
						ID: uuid.NewString(),
						MessageTo: domain.Profile{ID: "e2b5f92e-c31b-11eb-8529-0242ac130003"},
						MessageFrom: domain.Profile{ID: "424935b1-766c-4f99-b306-9263731518bc"},
						Path: "",
						Timestamp: time.Now(),
						Content: string(m.data),

					}*/

				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}

		}
	}
}
