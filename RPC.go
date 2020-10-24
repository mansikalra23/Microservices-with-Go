package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int //use to allevate functions to methods

var database []Item // slice of Item datatype

func (a *API) GetDB(empty string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database { //for loop to iterate values in database
		if val.Title == title {
			getItem = val
		}
	}
	*reply = getItem //error type
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error { //add an item to database and returns it
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error { //replcing the old value with new
	var changed Item

	for i, val := range database {
		if val.Title == edit.Title {
			database[i] = Item{edit.Title, edit.Body}
			changed = database[i]
		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error { // deleted the item from database
	var del Item
	for i, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:i], database[i+1:]...) //append spread opertor
			del = item
			break
		}

	}
	*reply = del
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("error registering: ", err)
	}

	log.Printf("serving rpc on port %d", 8080)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error registering: ", err)
	}

	//to generally print the database without client

	/*fmt.Println("Initial database: ", database)
	a := Item{"first", "a test item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("Database after adding: ", database)

	DeleteItem(b)
	fmt.Println("Database after deleting: ", database)

	EditItem("third", Item{"fourth", "a new item"})
	fmt.Println("database after editing: ", database)

	x := GetByName("fourth")
	y := GetByName("first")
	fmt.Println(x, y)*/
}
