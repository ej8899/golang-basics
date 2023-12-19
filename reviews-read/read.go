package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type fullName struct {
	fname string
	lname string
}

func main() {
	var fileName string
	var sliceNames []fullName

	fmt.Println("Enter filename:")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		splitName := strings.Split(fileScanner.Text(), " ")
		if len(splitName) == 2 {
			fname := splitName[0]
			lname := splitName[1]
			if len(fname) > 20 {
				fname = fname[:20]
			}
			if len(lname) > 20 {
				lname = lname[:20]
			}
			sliceNames = append(sliceNames, fullName{fname, lname})
		} else {
			fmt.Println("Invalid format")
		}
	}

	fmt.Println(sliceNames)

}
