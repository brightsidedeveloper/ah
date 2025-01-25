package handler

import (
	"log"
	"net/http"
	"server/internal/bin"
	"server/internal/buf"

	"google.golang.org/protobuf/proto"
)

type Handler struct {
	b *bin.Binary
}

func NewHandler(b *bin.Binary) *Handler {
	return &Handler{
		b: b,
	}
}

var users = []*buf.User{
	{
		Id:   1,
		Name: "Alice",
	},
	{
		Id:   2,
		Name: "Bob",
	},
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	data, err := proto.Marshal(&buf.Users{
		Users: users,
	})
	if err != nil {
		h.b.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}

	h.b.Respond(w, data, http.StatusOK)
}

func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := h.b.BytesFromBody(r.Body)
	if err != nil {
		h.b.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	var user buf.User
	if err := proto.Unmarshal(bodyBytes, &user); err != nil {
		http.Error(w, "Failed to decode user", http.StatusInternalServerError)
		log.Println("Error decoding user:", err)
		return
	}

	for _, u := range users {
		if u.Id == user.Id {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}
	}

	users = append(users, &user)

	data, err := proto.Marshal(&buf.Users{
		Users: users,
	})
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
