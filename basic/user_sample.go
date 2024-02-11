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
		Emails: []string{
			"sololeveling@gmail.com",
			"frieren@gmail.com",
		}, // repeated
		Gender: basic.Gender_GENDER_MALE, // enum
	}

	log.Println(&u)
}
