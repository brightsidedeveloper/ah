package bin

import (
	"io"
	"net/http"
	"server/internal/buf"

	"google.golang.org/protobuf/proto"
)

type Binary struct {
}

func NewBinary() *Binary {
	return &Binary{}

}

func (r *Binary) Respond(w http.ResponseWriter, binary []byte, status int) {
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.WriteHeader(status)
	w.Write(binary)
}

func (r *Binary) Error(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.WriteHeader(status)
	data, err := proto.Marshal(&buf.Error{
		Message: message,
	})
	if err != nil {
		w.Write([]byte("Failed to encode error message"))
		return
	}
	w.Write(data)
}

func (r *Binary) BytesFromBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	return io.ReadAll(body)
}
