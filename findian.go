package main

import (
	"fmt"
	"strings"
)

func main() {
	// Prompt user for input
	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)

	// Convert the input to lowercase for case-insensitive comparison
	lowercaseInput := strings.ToLower(input)

	// Check conditions and print the result
	if strings.HasPrefix(lowercaseInput, "i") &&
		strings.Contains(lowercaseInput, "a") &&
		strings.HasSuffix(lowercaseInput, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
