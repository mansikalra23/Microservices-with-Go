package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

// homepage function handle all requests to our root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is homepage.")
}

// handleRequests function match the URL path hit with a defined function
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// main function will kick off the API
func main() {
	handleRequests()
}
