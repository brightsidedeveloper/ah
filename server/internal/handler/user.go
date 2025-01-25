package handler

import (
	"net/http"
	"server/internal/buf"

	"google.golang.org/protobuf/proto"
)

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
	h.b.ProtoRespond(w, http.StatusOK, &buf.Users{
		Users: users,
	})
}

func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := h.b.BytesFromBody(r.Body)
	if err != nil {
		h.b.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	var user buf.User
	if err := proto.Unmarshal(bodyBytes, &user); err != nil {
		h.b.Error(w, "Failed to decode body", http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Id == user.Id {
			h.b.Error(w, "User already exists", http.StatusBadRequest)
			return
		}
	}

	users = append(users, &user)

	h.b.ProtoRespond(w, http.StatusCreated, &buf.Users{
		Users: users,
	})
}
