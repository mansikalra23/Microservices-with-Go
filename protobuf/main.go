package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	old := &Person{
		Name: "Mansi",
		Age:  19,
	}

	data, err := proto.Marshal(old)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// unmarshalling the data in new variable
	new := &Person{}
	err = proto.Unmarshal(data, new)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(new.GetAge(), new.GetName())
	// go run main.go person.pb.go
}
