package routes

import (
	"server/internal/handler"

	"github.com/go-chi/chi/v5"
)

func MountRoutes(r chi.Router, h *handler.Handler) {
	r.Get("/users", h.GetUsers)
	r.Post("/user", h.PostUser)
}
