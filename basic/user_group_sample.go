package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	budi := basic.User{
		Id:       1,
		Username: "budi",
		IsActive: true,
		Password: []byte("budipassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	joko := basic.User{
		Id:       2,
		Username: "joko",
		IsActive: true,
		Password: []byte("jokopassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	Intan := basic.User{
		Id:       3,
		Username: "intan",
		IsActive: true,
		Password: []byte("intanpassword"),
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	// Create Group
	developers := basic.UserGroup{
		GroupId:   1,
		GroupName: "Developers",
		Users: []*basic.User{
			&budi,
			&joko,
			&Intan,
		},
		Description: "Group for developers",
	}

	// Print Group
	jsonBytes, _ := protojson.Marshal(&developers)
	log.Println(string(jsonBytes))
}
