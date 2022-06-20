package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1] //0th arg is the go file name, so 1st arg is the name
	age := os.Args[2]

	var greeting string = createGreet(name, age)
	fmt.Printf("%s", greeting)

}

func createGreet(name string, age string) string {
	greeting := "Hello " + name + " :) " + "You are " + age + " years old"
	return greeting
}
