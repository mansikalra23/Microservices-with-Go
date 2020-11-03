// using gorilla mux router
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"' // json:"path`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	// entering the dummy data
	articles := Articles{
		Article{Title: "KloudOne", Desc: "Trainee", Content: "Welcome"},
	}

	fmt.Println("Endpoint hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles) // encodes articles array into json string
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post Endpoint Worked!")
}

// homepage function handle all requests to our root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is homepage.")
}

// handleRequests function match the URL path hit with a defined function
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	// myRouter lets us call same page wuth different methods that perform different functions
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// main function will kick off the API
func main() {
	handleRequests()
}
