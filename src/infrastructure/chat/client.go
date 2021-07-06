package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Hub Hub
}

func NewClient(hub Hub) Client {
	return Client{
		Hub: hub,
	}
}
func (c *Client) ServeWs(w http.ResponseWriter, r *http.Request, roomId string) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	fmt.Println(roomId)

	conn := &connection{send: make(chan []byte, 256), ws: ws}
	s := Subscription{conn, roomId, c.Hub}
	c.Hub.register <- s
	go s.WritePump()
	go s.ReadPump()
	return nil
}
