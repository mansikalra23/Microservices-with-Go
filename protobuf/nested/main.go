package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	old := Person{
		Name: "Mansi",
		Age:  19,
		Contact: &Contact{
			Number: 9876543210,
			Email:  "mansi@gmail.com",
		},
	}

	data, err := proto.Marshal(&old)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	new := &Person{}
	err = proto.Unmarshal(data, new)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(new.GetName())
	fmt.Println(new.GetAge())
	fmt.Println(new.Contact.GetNumber())
	fmt.Println(new.Contact.GetEmail())

}
