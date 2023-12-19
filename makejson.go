package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Print("Enter a name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Enter an address: ")
	var address string
	fmt.Scanln(&address)

	personMap := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println("JSON Object:")
	fmt.Println(string(jsonData))
}
