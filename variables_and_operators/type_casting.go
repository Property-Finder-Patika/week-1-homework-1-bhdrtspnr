package main

import "fmt"

//type conversions

func main() {
	/*
		var floatValue float32 = 5.45

		// type conversion from float to int
		var intValue int = int(floatValue)

		fmt.Printf("Float Value is %g\n", floatValue)
		fmt.Printf("Integer Value is %d", intValue)
	*/

	/*
		var intValue int = 2

		// type conversion from int to float
		var floatValue float32 = float32(intValue)

		fmt.Printf("Integer Value is %d\n", intValue)
		fmt.Printf("Float Value is %f", floatValue)
	*/

	var number1 int = 20
	var number2 float32 = 5.7
	var sum float32

	// addition of different data types
	sum = float32(number1) + number2

	fmt.Printf("Sum is %g", sum)
}
