package main

import (
	"encoding/csv"
	"io"
	"os"
)

type person struct {
	name  string
	email string
	date  string
}

func main() {

	fname, err := os.Open("data/data.csv")
	peopleFromFile := []person{}

	if err != nil {
		panic(err)
	}

	defer fname.Close()
	read := csv.NewReader(fname)

	final, err := os.Create("data/final.csv")

	if err != nil {
		panic(err)
	}

	defer final.Close()

	wr := csv.NewWriter(final)

	defer wr.Flush()
	headers := []string{"name", "email", "date"}
	wr.Write(headers)

	for {
		line, err := read.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		peopleFromFile = append(peopleFromFile, person{line[0], line[1], line[2]})
	}

	for _, row := range peopleFromFile {
		wr.Write([]string{row.name, row.email, row.date})
	}

}
