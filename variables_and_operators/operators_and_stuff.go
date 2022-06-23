package main

//operators, increment, decrement, bitwise shifting, basic pointers

import (
	"fmt"
)

func main() {

	number1 := 10
	number2 := 3

	sum := number1 + number2
	subtract := number1 - number2
	multiply := number1 * number2
	divide := number1 / number2
	modulo := number1 % number2

	fmt.Printf("Total is: %d \n", sum)
	fmt.Printf("Total is: %d \n", subtract)
	fmt.Printf("Total is: %d \n", multiply)
	fmt.Printf("Total is: %d \n", divide)
	fmt.Printf("Total is: %d", modulo)

	//Increment and decrement by one
	//number1++
	//number1--

	isTrue := number1 == number2
	fmt.Printf("is true? %t", isTrue)

	num := 212

	/*	212 = 11010100 (In binary)
		212 >> 3 = 00011010 = 26 (Right shift by 3 bits)
		212 >> 7 = 00000001 = 1 (Right shift by 7 bits)
		212 >> 0 = 11010100 = 212 (No Shift)

	*/
	for i := 0; i <= 3; i++ {
		fmt.Printf("Right Shift by %d: %d\n", i, num>>i)
	}

	/*
		212 = 11010100 (In binary)
		212 << 1 = 110101000 = 424 (Adds one 0 to the right)
		212 << 3 = 11010100000 = 1696 (Adds three 0's to the right)
		212 << 0 = 11010100 = 212 (No Shift)
	*/
	for i := 0; i <= 3; i++ {

		// left shift 212 by bits from 0 to 3
		fmt.Printf("Left Shift by %d: %d\n", i, num<<i)

	}

	// &num prints the
	// memory address where 20 is stored
	fmt.Println(&num) //0xc000012098 virtual memory address

	// pointer declaration
	b := 20
	var num1 *int = &b //num1 pretty much points to b's location on the memory

	// gives the memory address
	fmt.Println(num1)

}
