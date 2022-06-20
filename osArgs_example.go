package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1] //0th arg is the go file name, so 1st arg is the name
	var greeting string = createGreet(name)
	fmt.Printf("%s", greeting)

}

func createGreet(name string) string {
	greeting := "Hello " + name + " :) "
	return greeting
}
