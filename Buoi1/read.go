package main

import (
	"fmt"
	"os"
)

type name struct {
	fname string
	lname string
}

func main() {
	// Prompt user for the name of the text file
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scan(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read each line and create a struct for each name
	var names []name
	var fname, lname string
	for {
		_, err := fmt.Fscanf(file, "%s %s\n", &fname, &lname)
		if err != nil {
			break
		}
		names = append(names, name{fname, lname})
	}

	// Print the first and last names for each struct
	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.fname, n.lname)
	}
}
