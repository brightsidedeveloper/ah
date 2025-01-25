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

func (b *Bin) Write(w http.ResponseWriter, status int, binary []byte) {
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.WriteHeader(status)
	w.Write(binary)
}

func (b *Bin) WriteError(w http.ResponseWriter, status int, message string) {
	data, err := proto.Marshal(&buf.Error{
		Message: message,
	})
	if err != nil {
		w.Write([]byte("Failed to encode error message"))
		return
	}
	b.Write(w, status, data)
}

func (b *Bin) bytesFromBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	return io.ReadAll(body)
}

func (b *Bin) UnmarshalBody(body io.ReadCloser, protoMessage proto.Message) error {
	data, err := b.bytesFromBody(body)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, protoMessage)
}

func (b *Bin) ProtoWrite(w http.ResponseWriter, status int, protoMessage proto.Message) {
	data, err := proto.Marshal(protoMessage)
	if err != nil {
		b.WriteError(w, http.StatusInternalServerError, "Failed to encode message")
		return
	}
	b.Write(w, status, data)
}
