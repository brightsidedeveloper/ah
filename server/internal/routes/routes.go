package routes

import (
	"server/internal/handler"

	"github.com/go-chi/chi/v5"
)

func MountRoutes(r chi.Router, h *handler.Handler) {

	api := chi.NewRouter()
	api.Get("/users", h.GetUsers)
	api.Post("/user", h.PostUser)

	r.Mount("/api", api)
}
