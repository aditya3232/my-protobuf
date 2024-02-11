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
	addr := basic.Address{
		Street:  "Jl. Kebon Sirih",
		City:    "Jakarta",
		Country: "Indonesia",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  6.174465,
			Longitude: 106.822745,
		},
	}

	p := basic.User{
		Id:       99,
		Username: "nisa",
		IsActive: true,
		Password: []byte("nisapassword"),
		Emails: []string{
			"kusuriyanohitorigoto@gmail.com",
			"mashle@gmail.com",
		}, // repeated
		Gender:  basic.Gender_GENDER_FEMALE, // enum
		Address: &addr,                      // embedded
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}
