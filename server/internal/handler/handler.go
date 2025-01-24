package handler

import (
	"io"
	"log"
	"net/http"
	"server/internal/buf"

	"google.golang.org/protobuf/proto"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
	w.Header().Set("Content-Type", "application/x-protobuf")

	data, err := proto.Marshal(&buf.Users{
		Users: users,
	})

	if err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-protobuf")

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		log.Println("Error reading body:", err)
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
