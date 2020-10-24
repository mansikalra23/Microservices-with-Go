package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item //make shift database

	client, err := rpc.DialHTTP("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}
	c := Item{"Third", "A third item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db) //to get database with values

	fmt.Println("Database after adding: ", db)

	client.Call("API.EditItem", Item{"Second", "A new second item"}, &reply)
	fmt.Println("Database after editing: ", db)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database after deleting: ", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item: ", reply)
}
