package handler

import (
	"server/internal/bin"
)

type Handler struct {
	bin *bin.Bin
}

func NewHandler(b *bin.Bin) *Handler {
	return &Handler{
		bin: b,
	}
}
