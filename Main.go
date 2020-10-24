package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye KloudOne!")
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello KloudOne!")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error!", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", b)
	})

	log.Println("Starting Server")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
