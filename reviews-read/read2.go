package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	maxLength = 20
)

type Name struct {
	fname string
	lname string
}

// func (n *Name) Set(firstname string, lastname string) {
// 	n.fname = firstname
// 	n.lname = lastname

// 	if len(firstname) > maxLength {
// 		n.fname = firstname[:maxLength]
// 	}
// 	if len(lastname) > maxLength {
// 		n.lname = lastname[:maxLength]
// 	}
// }

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	names := make([]Name, 0, 3)
	fmt.Println("Welcome, please enter the name of the text file: ")
	filename, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	filename = strings.Trim(filename, "\n")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// works like destructor
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var person Name
		if len(s[0]) > 20 {
			person.fname = s[0][:20]
			if len(s[1]) > 20 {
				person.lname = s[1][:20]
			}
		} else if len(s[1]) > 20 {
			person.fname = s[0]
			person.lname = s[1][:20]
		} else {
			person.fname, person.lname = s[0], s[1]
		}
		names = append(names, person)
	}

	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}

}