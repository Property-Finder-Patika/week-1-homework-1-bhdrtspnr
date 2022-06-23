package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var Magazines []Magazine //our db in most simple terms

type Magazine struct { //create magazine struct
	Id      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func addMagazine(title string, desc string, content string) { //for local usage NOT FOR THE API
	var newMagazine Magazine = Magazine{Title: title, Desc: desc, Content: content}
	Magazines = append(Magazines, newMagazine)
}

func AppendTestMagazines() { //test data to check if our app is working
	Magazines = append(Magazines, Magazine{Id: "1", Title: "Test", Desc: "Test", Content: "Test"})
	Magazines = append(Magazines, Magazine{Id: "2", Title: "Test2", Desc: "Test2", Content: "Test2"})
}

func returnAllMagazines(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Magazines) //returns the magazines slice as json
}

func retrieveMagazine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: retrieveMagazine")
	for i := 0; i < len(Magazines); i++ { //check if we have a magazine with given key as id
		if Magazines[i].Id == key {
			json.NewEncoder(w).Encode(Magazines[i]) //return the magazine with given key as id
		}
	}
}

func createNewMagazine(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body) //get request
	var magazine Magazine
	json.Unmarshal(reqBody, &magazine)      //parse the request as our magazine struct
	Magazines = append(Magazines, magazine) //append it to our db (slice)
	json.NewEncoder(w).Encode(magazine)     //return the added object in response
}

func homePage(w http.ResponseWriter, r *http.Request) { //simple home page
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func deleteMagazine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //parse data from request
	key := vars["id"]   //get document key

	//more of a pythonic iteration, i for index, item for data stored at index i
	//saw here: https://stackoverflow.com/questions/41445216/go-loop-indices-for-range-on-slice
	for i, item := range Magazines { //check if we have a magazine with given key as id
		if item.Id == key {
			Magazines = append(Magazines[:i], Magazines[i+1:]...) //delete the magazine with given key as id and update the so called db
		}
	}

}
