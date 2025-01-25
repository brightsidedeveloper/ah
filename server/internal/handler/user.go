package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/buf"
	"strconv"
	"time"

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

type OldUser struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

var oldUsers = []OldUser{
	{
		Id:   1,
		Name: "Alice",
	},
	{
		Id:   2,
		Name: "Bob",
	},
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	// Measure JSON serialization time
	startJSON := time.Now()
	jsonData, err := json.Marshal(oldUsers)
	if err != nil {
		http.Error(w, "Failed to encode users as JSON", http.StatusInternalServerError)
		return
	}
	jsonDuration := time.Since(startJSON)

	// Measure Protobuf serialization time
	startProto := time.Now()
	protoData, err := proto.Marshal(&buf.Users{
		Users: users,
	})
	if err != nil {
		h.b.Error(w, http.StatusInternalServerError, "Failed to encode users as Protobuf")
		return
	}
	protoDuration := time.Since(startProto)

	// Log the comparison
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"json_size":      strconv.FormatInt(int64(len(jsonData)), 10),
		"json_duration":  jsonDuration.String(),
		"proto_size":     strconv.FormatInt(int64(len(protoData)), 10),
		"proto_duration": protoDuration.String(),
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
