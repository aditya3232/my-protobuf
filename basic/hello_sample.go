package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

// BasicHello is a function that demonstrates the use of the generated code
func BasicHello() {
	h := basic.Hello{
		Name: "Hello, Aditya!",
	}

	log.Println(&h)
}
