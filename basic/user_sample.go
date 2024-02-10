package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
	u := basic.User{
		Id:       99,
		Username: "aditya",
		IsActive: true,
		Password: []byte("adityapassword"),
	}

	log.Println(&u)
}
