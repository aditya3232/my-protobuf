package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
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

	comm := randomCommunicationChannel()

	p := basic.User{
		Id:       99,
		Username: "nisa",
		IsActive: true,
		Password: []byte("nisapassword"),
		Emails: []string{
			"kusuriyanohitorigoto@gmail.com",
			"mashle@gmail.com",
		}, // repeated
		Gender:               basic.Gender_GENDER_FEMALE, // enum
		Address:              &addr,
		CommunicationChannel: comm, // random message
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

// func for random communication channel
func randomCommunicationChannel() *anypb.Any {
	paperMail := &basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := &basic.SocialMedia{
		SocialMediaPlatform: "byteMe",
		SocialMediaUsername: "krypton.man",
	}

	instantMessaging := &basic.InstantMessaging{
		InstantMessagingProduct:  "whatsup",
		InstantMessagingUsername: "krypton.last",
	}

	var a *anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		any, _ := anypb.New(paperMail)
		a = any
	case 1:
		any, _ := anypb.New(socialMedia)
		a = any
	default:
		any, _ := anypb.New(instantMessaging)
		a = any
	}

	return a
}

// unmarshal any known type
func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (Social Media)
	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

// unmarshal any not known type from random communication channel
func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()

	if err != nil {
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

// MessageIs untuk mengatahui tipe message yang ada di any
// disini akan di cek apakah message yang ada di any adalah PaperMail
// jika iya, maka akan di unmarshal ke PaperMail
func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, but :", a.TypeUrl)
	}
}
