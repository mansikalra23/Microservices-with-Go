package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Trainee - Structure for trainee information
type Trainee struct {
	Name  string `json:"Name"`
	Batch string `json:"Batch"`
}

var Trainees []Trainee

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to KloudOne!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllTrainees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTrainees")
	json.NewEncoder(w).Encode(Trainees)
}

func returnSingleTrainee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Name"]

	for _, trainee := range Trainees {
		if trainee.Name == key {
			json.NewEncoder(w).Encode(trainee)
		}
	}
}

func createNewTrainee(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Trainee struct
	// append this to our Trainees array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var trainee Trainee
	json.Unmarshal(reqBody, &trainee)
	// update our global Trainees array to include
	// our new Trainee
	Trainees = append(Trainees, trainee)

	json.NewEncoder(w).Encode(trainee)
}

func deleteTrainee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for index, trainee := range Trainees {
		if trainee.Name == name {
			Trainees = append(Trainees[:index], Trainees[index+1:]...)
		}
	}
	fmt.Println("Endpoint Hit: deleteTrainee")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/trainees", returnAllTrainees)
	myRouter.HandleFunc("/trainee", createNewTrainee).Methods("POST")
	myRouter.HandleFunc("/trainee/{name}", deleteTrainee).Methods("DELETE")
	myRouter.HandleFunc("/trainee/{name}", returnSingleTrainee)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Trainees = []Trainee{
		Trainee{Name: "abc", Batch: "A"},
		Trainee{Name: "efg", Batch: "B"},
	}
	handleRequests()
}
