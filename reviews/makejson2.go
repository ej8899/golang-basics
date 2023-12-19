package main
import (
"fmt"
"os"
"encoding/json"
"bufio"
"strings"
)

func main(){

	personMap := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter the name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Please enter the address: ")
	address, _ := reader.ReadString('\n')
	personMap["name"] = strings.TrimSuffix(name, "\r\n")
	personMap["address"] =  strings.TrimSuffix(address, "\r\n")
	jsonStructure, _ := json.Marshal(personMap)
    fmt.Println(string(jsonStructure))


}