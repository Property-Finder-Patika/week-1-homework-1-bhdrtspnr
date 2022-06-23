package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Please enter your name")
	fmt.Scanln(&name)

	var greeting = createGreet(name) //assignment of the return value of the function to a variable
	fmt.Printf("%s", greeting)

}

func createGreet(name string) string {
	greeting := "Hello " + name + " :) "
	return greeting

}
