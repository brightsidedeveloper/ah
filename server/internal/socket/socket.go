package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	id   string
	name string
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		id:   conn.RemoteAddr().String(),
	}
}

func (c *Client) ReadMessages(s *Server) {
	defer func() {
		s.mut.Lock()
		delete(s.clients, c.id)
		s.mut.Unlock()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
