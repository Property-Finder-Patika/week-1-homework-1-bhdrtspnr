package main

import (
	"fmt"
)

func main() {

	fmt.Print(createGreet("Bahadir"))
}

func createGreet(name string) string { //a function takes a string argument "name" and returns a string
	greeting := "Hello " + name + " :) " //declaration and assignment at the same time with :=
	return greeting
}
