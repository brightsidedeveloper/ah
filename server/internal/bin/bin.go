package bin

import (
	"io"
	"net/http"
	"server/internal/buf"

	"google.golang.org/protobuf/proto"
)

type Bin struct {
}

func NewBinary() *Bin {
	return &Bin{}

}

func (r *Bin) Respond(w http.ResponseWriter, status int, binary []byte) {
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.WriteHeader(status)
	w.Write(binary)
}

func (r *Bin) Error(w http.ResponseWriter, status int, message string) {
	data, err := proto.Marshal(&buf.Error{
		Message: message,
	})
	if err != nil {
		w.Write([]byte("Failed to encode error message"))
		return
	}
	r.Respond(w, status, data)
}

func (r *Bin) BytesFromBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	return io.ReadAll(body)
}

func (r *Bin) ProtoRespond(w http.ResponseWriter, status int, protoMessage proto.Message) {
	data, err := proto.Marshal(protoMessage)
	if err != nil {
		r.Error(w, http.StatusInternalServerError, "Failed to encode message")
		return
	}
	r.Respond(w, status, data)
}
