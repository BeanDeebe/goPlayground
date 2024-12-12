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
	date  string
}

func main() {

	data, err := os.Open("data/data.csv")
	list_of_people := []person{}

	if err != nil {
		panic(err)
	}

	defer data.Close()
	reader := csv.NewReader(data)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		currName, currEmail, currDate := line[0], line[1], line[2]

		currPerson := person{currName, currEmail, currDate}
		list_of_people = append(list_of_people, currPerson)

	}

	fmt.Print(list_of_people)

}
