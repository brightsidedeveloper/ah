package handler

import (
	"server/internal/bin"
	"server/internal/query"
)

type Handler struct {
	bin   *bin.Bin
	query *query.Queries
}

func NewHandler(b *bin.Bin, q *query.Queries) *Handler {
	return &Handler{
		bin:   b,
		query: q,
	}
}
