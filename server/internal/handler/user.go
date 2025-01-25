package handler

import (
	"fmt"
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
		h.b.Error(w, http.StatusBadRequest, "Failed to read body")
		return
	}

	var user buf.User
	if err := proto.Unmarshal(bodyBytes, &user); err != nil {
		h.b.Error(w, http.StatusBadRequest, "Failed to decode body")
		return
	}

	if err := createUser(&user); err != nil {
		h.b.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	h.b.ProtoRespond(w, http.StatusCreated, &buf.Users{
		Users: users,
	})
}

func createUser(user *buf.User) error {
	for _, u := range users {
		if u.Id == user.Id {
			return fmt.Errorf("user with id %d already exists", user.Id)
		}
	}
	users = append(users, user)
	return nil
}
