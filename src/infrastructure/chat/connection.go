package chat

import (
	"github.com/gorilla/websocket"
	"time"
)

type connection struct {
	ws *websocket.Conn
	send chan []byte
	connectionId string
}

func (c *connection) write(mt int, payload []byte) error {
	if err := c.ws.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		return err
	}
	return c.ws.WriteMessage(mt, payload)
}

