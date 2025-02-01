package handler

import (
	"net/http"
	"server/internal/socket"
)

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := socket.NewClient(conn)

	go c.ReadMessages(h.ss)
}
