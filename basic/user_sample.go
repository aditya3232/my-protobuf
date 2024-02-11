package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
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

// protobuf to json
func ProtoToJsonUser() {
	p := basic.User{
		Id:       99,
		Username: "nisa",
		IsActive: true,
		Password: []byte("nisapassword"),
		Emails: []string{
			"kusuriyanohitorigoto@gmail.com",
			"mashle@gmail.com",
		}, // repeated
		Gender: basic.Gender_GENDER_FEMALE, // enum
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}
