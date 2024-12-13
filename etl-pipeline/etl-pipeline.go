package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type person struct {
	name  string
	email string
	date  string // could update this to use a time package?
}

func main() {

	fname, err := os.Open("data/data.csv")

	peopleFromFile := []person{}

	if err != nil {
		panic(err)
	}

	defer fname.Close()

	read := csv.NewReader(fname)

	for {
		line, err := read.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		personName, personEmail, personDate := line[0], line[1], line[2]

		newPerson := person{personName, personEmail, personDate}
		peopleFromFile = append(peopleFromFile, newPerson)
	}

	for i := 0; i < len(peopleFromFile); i++ {
		fmt.Println(peopleFromFile[i])
	}
}
