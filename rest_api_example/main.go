package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	// creates a new instance of a mux router using this helps instead of using the default http, says someone on the internet
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllMagazines)
	myRouter.HandleFunc("/magazine/{id}", retrieveMagazine).Methods("GET") //unless specified as GET, it overrides the delete method
	myRouter.HandleFunc("/magazine", createNewMagazine).Methods("POST")
	myRouter.HandleFunc("/magazine/{id}", deleteMagazine).Methods("DELETE") //only allow DELETE requests -> same endpoint, creating another method

	log.Fatal(http.ListenAndServe(":10000", myRouter)) //listen to 10000 port
}

var testing bool = true

func main() {
	if testing {
		AppendTestMagazines() //add test magazines
	}
	handleRequests()
}
