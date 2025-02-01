package handler

import (
	"server/internal/bin"
	"server/internal/query"
	"server/internal/socket"
)

type Handler struct {
	bin   *bin.Bin
	query *query.Queries
	ss    *socket.Server
}

func NewHandler(b *bin.Bin, q *query.Queries, s *socket.Server) *Handler {
	return &Handler{
		bin:   b,
		query: q,
		ss:    s,
	}
}
