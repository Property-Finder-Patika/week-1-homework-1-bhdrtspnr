package main

import (
	"fmt"
	"strconv"
)

var persons []Person //person slice to store the created persons (pretty much person db)

type Person struct { //struct for persons, looks pretty familiar to java objectss
	name   string
	age    int
	height int
	weight int
}

func (p Person) toString() string { //not sure if there's already a to string function
	return "Name: " + p.name + "\nAge: " + strconv.Itoa(p.age) + "\nHeight: " + strconv.Itoa(p.height) + "\nWeight: " + strconv.Itoa(p.weight) //converting integers to strings, not sure if there's an easier way to do this ??
}

func addPerson(name string, age int, height int, weight int) { //adds a person to the slice persons slice
	var newPerson Person = Person{name: name, age: age, height: height, weight: weight}
	persons = append(persons, newPerson)
}

func listPeople() { //iterate over list to print all people structss
	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i].toString())
	}
}

func searchPerson() {
	fmt.Println("Enter the name of the person to search")
	var name string
	fmt.Scanln(&name)
	for i := 0; i < len(persons); i++ { //iterate over the slice to find the person
		if persons[i].name == name { //check any of the names is matching the given argument
			fmt.Println(persons[i].toString())
		}
	}
}

func deletePerson() {
	fmt.Println("Enter the name of the person to delete")
	var name string
	fmt.Scanln(&name)
	for i := 0; i < len(persons); i++ { //check any of the names is matching the given argument
		if persons[i].name == name {
			persons = append(persons[:i], persons[i+1:]...) //not sure how to delete from a slice but someone on the stackoverlow suggested this https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
		}
	}
}

func isPersonExists(name string) bool {
	for i := 0; i < len(persons); i++ {
		if persons[i].name == name {
			return true
		}
	}
	return false
}
