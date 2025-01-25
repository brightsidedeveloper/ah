package handler

import (
	"encoding/json"
	"net/http"
	"server/internal/buf"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

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
		h.bin.WriteError(w, http.StatusInternalServerError, "Failed to encode users as Protobuf")
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
