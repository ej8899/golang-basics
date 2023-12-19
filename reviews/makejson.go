package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var inputName string
	var inputAddress string
	nameAddressMap := make(map[string]string)

	fmt.Println("Enter a name:")
	fmt.Scan(&inputName)
	fmt.Println("Enter a address:")
	fmt.Scan(&inputAddress)
	nameAddressMap[inputName] = inputAddress

	nameAddressJson, err := json.Marshal(nameAddressMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(nameAddressJson))
}
