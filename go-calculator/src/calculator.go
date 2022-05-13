// Authored by		: Markus Walker
// Date Modified	: 5/13/22

// Description		: Simple calculator script written in Go to perform the following operations:
//					  addition, subtraction, multiplication and division.

package main

import (
	"fmt"
	"os"

	"github.com/pborman/getopt"
)

// Function to add the two numbers.
func Addition(a, b int) int {
	return a + b
}

// Function to subtract the two numbers.
func Subtraction(a, b int) int {
	return a - b
}

// Function to multiply the two numbers.
func Multiplication(a, b int) int {
	return a * b
}

// Function to divide the two numbers.
func Division(a, b int) int {
	return a / b
}

// Function to display the usage of the program.
func usage() {
	optHelp := getopt.BoolLong("help", 0, "Display usage help")
	getopt.Parse()

	// If user views the `--help` flag, then show the usage and exit out of the program.
	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
}

// Driver function.
func main() {
	usage()
	var num1, num2, choice int

	fmt.Println("\x1B[96m===========================================")
	fmt.Println("\tGolang Calculator Tool")
	fmt.Println("===========================================\x1B[0m")

	fmt.Println("\nWelcome to the Golang Calculator Tool! Please find the supported operations:")
	fmt.Print("\n 1: Addition")
	fmt.Print("\n 2: Subtraction")
	fmt.Print("\n 3: Multiplication")
	fmt.Print("\n 4: Division\n")

	// Prompt the user for the choice that they wish to do.
	fmt.Print("\nEnter an operation that you wish to do:")
	fmt.Scan(&choice)

	// Prompt the user for input for the first and second number.
	fmt.Print("Enter the first number:")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number:")
	fmt.Scan(&num2)

	// Perform operation based upon what the use choes to do. By default, exit out of the program if they don't pick a correct value...
	switch choice {
	case 1:
		result := Addition(num1, num2)
		fmt.Println(num1, " + ", num2, " = ", result)
	case 2:
		result := Subtraction(num1, num2)
		fmt.Println(num1, " - ", num2, " = ", result)
	case 3:
		result := Multiplication(num1, num2)
		fmt.Println(num1, " * ", num2, " = ", result)
	case 4:
		result := Division(num1, num2)
		fmt.Println(num1, " / ", num2, " = ", result)
	default:
		fmt.Println("Invalid choice. Please choose betwen 1-4 and try again.")
	}
}
