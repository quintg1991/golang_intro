// File name: blueberry.go
// Description: blueberry program implementation of the blueberry program
//
// Version: 1.1
// Date: 4/12/19
// Author: Franklin Glover

package main

import (
	// Inport I/O functions like print and scan
	"fmt"
)

func main() {

	// Declare variables
	var first string
	var second string
	var iterations int

	// Get first input
	fmt.Print("Enter the first word: ")
	fmt.Scanln(&first)

	// Get second input
	fmt.Print("Enter the second word: ")
	fmt.Scanln(&second)

	// Get iterations input
	fmt.Print("Enter the number of iterations: ")
	fmt.Scanln(&iterations)

	// Let's take a look at a for loop
	for i := 1; i < iterations+1; i++ {
		if i%15 == 0 {
			fmt.Println(first, second)
			// Pay attention to how the syntax is specifically written here
		} else if i%5 == 0 {
			fmt.Println(second)
		} else if i%3 == 0 {
			fmt.Println(first)
		} else {
			fmt.Println(i)
		}
	}
}
