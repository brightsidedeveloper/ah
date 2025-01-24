package main

import (
	"fmt"
	"log"
	"server/internal/buf"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &buf.User{
		Id:    5,
		Name:  "John",
		Email: "",
	}
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	var newUser buf.User
	if err := proto.Unmarshal(data, &newUser); err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(map[int32]string{
		newUser.Id: newUser.Name,
	})
}
