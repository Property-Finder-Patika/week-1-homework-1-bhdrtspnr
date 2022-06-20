package main

import (
	"fmt"
	"strconv"
)

var schools []School

type School struct {
	Name     string
	City     string
	Students []Person
}

func (s School) toString() { //not returning string as it was in the person this time
	fmt.Println("Name: " + s.Name)
	fmt.Println("Cİty:" + s.City)
	fmt.Println("Student count: " + strconv.Itoa(len(s.Students))) //count of the students is the length of the students slice for each school

	for i := 0; i < len(s.Students); i++ { //iterate over the students to print all of them
		fmt.Println(s.Students[i].toString())
	}
}

func addSchool(name string, city string) { //adds a school to the slice schools
	var newSchool School = School{Name: name, City: city}
	schools = append(schools, newSchool)
}

func addStudent() { //not finished yet
	/*
		fmt.Println("Enter the name of the school to add the student to")
		var schoolName string
		fmt.Scanln(&schoolName)
		fmt.Println("Enter the name of the student")
		var studentName string
		fmt.Scanln(&studentName)

		for i := 0; i < len(schools); i++ { //iterate over the schools to find the school with the given name
			if schools[i].Name == schoolName {
				schools[i].Students = append(schools[i].Students, persons{Name: studentName}) //add the student to the school's students slice
			}
		}

	*/
}

func listSchools() { //iterate over list to print all schools
	for i := 0; i < len(schools); i++ {
		schools[i].toString()
	}
}
func searchSchool() { //iterate over the schools to find the school with the given name
	fmt.Print("Enter the name of the school to search: ")
	var schoolName string
	fmt.Scanln(&schoolName)
	for i := 0; i < len(schools); i++ {
		if schools[i].Name == schoolName {
			schools[i].toString()
		}
	}
}

func deleteSchool() {
	fmt.Print("Enter the name of the school to delete: ")
	var schoolName string
	fmt.Scanln(&schoolName)
	for i := 0; i < len(schools); i++ {
		if schools[i].Name == schoolName {
			schools = append(schools[:i], schools[i+1:]...) //same deletion used in the person struct
		}
	}
}
