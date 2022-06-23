package main

//use go mod init on terminal before using go run . else it won't work
//practicing structs, loops, slices, functions, and text formattingg

import (
	"fmt"
	"os"
)

var testing bool = false

func main() {

	if testing {
		fmt.Println("Testing stuff")
		var examplePerson Person = Person{name: "Bahadir", age: 25, height: 175, weight: 75}
		fmt.Print(examplePerson.toString())
	}

	var menuFlag bool = true

	for menuFlag {
		mainMenu()
	}

}

func mainMenu() {
	fmt.Println("")
	fmt.Println("Press 1 to enter the person menu")
	fmt.Println("Press 2 to enter the school menu")
	var menuInput int
	fmt.Scanln(&menuInput)
	switch menuInput {
	case 1:
		personMenu()
	case 2:
		schoolMenu()
	}
}

func schoolMenu() { //function to call school menu over and over again
	fmt.Println("Enter 1 to add a new school")
	fmt.Println("Enter 2 to list all schools")
	fmt.Println("Enter 3 to search for a school")
	fmt.Println("Enter 4 to delete a school")
	fmt.Println("Enter 5 to add student to school")
	fmt.Println("Enter 6 to exit")
	var menuInput int
	fmt.Scanln(&menuInput)
	switch menuInput {
	case 1:
		{
			fmt.Println(("Enter the name of the school"))
			var name string
			fmt.Scanln(&name)
			fmt.Println(("Enter the city of the school"))
			var city string
			fmt.Scanln(&city)
			addSchool(name, city)
		}
	case 2:
		listSchools()
	case 3:
		searchSchool()
	case 4:
		deleteSchool()
	case 5:
		addStudent()
	case 6:
		os.Exit(0)

	}
}

func personMenu() { //same with school menu function
	fmt.Println("Enter 1 to add a new person")
	fmt.Println("Enter 2 to list all people")
	fmt.Println("Enter 3 to search for a person")
	fmt.Println("Enter 4 to delete a person")
	fmt.Println("Enter 5 to exit")
	var menuInput int
	fmt.Scanln(&menuInput)

	switch menuInput {
	case 1:
		addPersonMenu()
	case 2:
		listPeople()
	case 3:
		searchPerson()
	case 4:
		deletePerson()
	case 5:
		os.Exit(0)

	default:
		fmt.Println("Please use a valid entry")
	}
}

func addPersonMenu() {
	fmt.Println(("Enter the name of the person"))
	var name string
	fmt.Scanln(&name)
	fmt.Println(("Enter the age of the person"))
	var age int
	fmt.Scanln(&age)
	fmt.Println(("Enter the height of the person"))
	var height int
	fmt.Scanln(&height)
	fmt.Println(("Enter the weight of the person"))
	var weight int
	fmt.Scanln(&weight)
	addPerson(name, age, height, weight)

}
