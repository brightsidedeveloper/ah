package handler

import (
	"server/internal/bin"
)

type Handler struct {
	b *bin.Bin
}

func NewHandler(b *bin.Bin) *Handler {
	return &Handler{
		b: b,
	}
}
