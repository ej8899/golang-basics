package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the name struct
type Name struct {
	FirstName string
	LastName  string
}

func main() {
	var fileName = "data.txt"
	
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	var names []Name

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into first and last names
		nameParts := strings.Fields(line)

		// Create a Name struct and add it to the slice
		nameStruct := Name{
			FirstName: truncateString(nameParts[0], 20),
			LastName:  truncateString(nameParts[1], 20),
		}
		names = append(names, nameStruct)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Print the first and last names from each struct in the slice
	fmt.Println("Names from the file:")
	for _, name := range names {
		fmt.Printf("First Name: %-20s Last Name: %-20s\n", name.FirstName, name.LastName)
	}
}

// Helper function to truncate or pad a string to a specified length
func truncateString(s string, length int) string {
	if len(s) > length {
		return s[:length]
	}
	return s + strings.Repeat(" ", length-len(s))
}
