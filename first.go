package main

import (
	"fmt"
	"go-basics/helper"
	"go-basics/routines"
)

func main() {
	fmt.Println("GO !!! ")

	// Initialization of Variables
	var name string
	var age int

	// Print to console
	fmt.Println("Enter Your name")
	// Get User Input
	fmt.Scan(&name)
	fmt.Println("Enter Your age")
	fmt.Scan(&age)

	// Print type of variables
	fmt.Printf("\nName Type: %T \n Age type : %T", name, age)

	// Arrays
	var hobbies [15]string

	fmt.Printf("\nEnter your hobby: ")
	fmt.Scan(&hobbies[0])

	// Inline print
	fmt.Printf("\nHi %v, Hope you are enjoying your life at %v with %v", name, age, hobbies[0])

	// Slice
	var testSlice []string
	testSlice = append(testSlice, "no")
	fmt.Printf("\nSlice Values: %v", testSlice)

	// Variable short hand initialization
	shortVal := 10
	shortVal = 20
	fmt.Println(shortVal)

	//While loop

	var input int
	choice := 0
	fmt.Println("Press 1 to enter while loop. To skip press any number")
	fmt.Scan(&choice)
	if choice == 1 {
		for {
			fmt.Println("Press 1 to break and any key to continue loop")
			fmt.Scan(&input)
			if input == 1 {
				break
			}
		}
	}
	fmt.Printf("\nFactorial of 5 is %v", helper.ComputeFact(5))

	var testArray [4]string
	testArray[0] = "test"
	testArray[1] = "asf"
	testArray[2] = "gfgf"
	testArray[3] = "reoi"

	// Usual for each for iteration
	for index, value := range testArray {
		fmt.Printf("\n Index: %v , Value : %v", index, value)
	}

	// _ is an empty placeholder variable
	// index is not needed we can use this syntax
	for _, val := range testArray {
		fmt.Printf("\n val: %v", val)
	}

	fmt.Println("\n\n*********************\nGo Routine\n*********************")
	routines.UseRoutine()

}

// function with single return type
// return type defined after arguments
func computeValues(size int, values []string) []string {
	var result []string
	if size > 0 {
		for index, val := range values {
			result[index] = val + "test"
		}
	}
	return result
}

// Function with multiple return types
// Func parameters defined in first paranthesis and return types in second
func mutipleReturnValues(size int, values string) (int, string) {
	return size + 2, values + "a"
}
