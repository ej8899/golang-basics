package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BubbleSort performs the Bubble Sort algorithm on a slice of ints:
func BubbleSort(arr []int) {
	n := len(arr)

	// Iterate through each element in the slice
	for i := 0; i < n-1; i++ {
		// Last 'i' elements are already sorted, so we don't need to check them
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements and swap them if in wrong order
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

// Swap the contents of a slice at the specified index with the contents at index+1
func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

// convert the input string to a slice of integers
func processInput(input string) ([]int, error) {
	var numbers []int

	// Split the input by whitespace
	inputs := strings.Fields(input)

	// Convert each substring to an integer
	for _, strNum := range inputs {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func main() {
	fmt.Println("Enter up to 10 integers (separated by spaces):")

	// Use bufio.Scanner to read the entire line - why  is input so 'difficult' in go?
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Process input
	numbers, err := processInput(input)
	if err != nil {
		fmt.Println("Error processing input:", err)
		return
	}

	// Perform Bubble Sort
	BubbleSort(numbers)

	// Print the sorted result
	fmt.Println("Sorted result:", numbers)
}
