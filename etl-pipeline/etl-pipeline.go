package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type finalRow struct {
	firstName string
	lastName  string
	userName  string
	email     string
	domain    string
	date      string
}

func main() {

	fname, err := os.Open("data/data.csv")
	finalConstr := []finalRow{}

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
	headers := []string{"firstName", "lastName", "userName", "email", "domain", "date"}
	wr.Write(headers)

	for {
		line, err := read.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if line[0] == "name" {
			continue
		}

		finalConstr = append(finalConstr, processRow(line))
	}

	for _, row := range finalConstr {
		wr.Write([]string{row.firstName, row.lastName, row.userName, row.email, row.domain, row.date})
	}

}

func processName(arrName []string) []string {
	firstName := arrName[0]
	lastName := arrName[1]
	userArr := []string{firstName[:1], lastName[1:]}
	userName := strings.Join(userArr, "")
	userName = strings.ToLower(userName)

	finalArr := []string{firstName, lastName, userName}
	return finalArr
}

func processRow(arrRow []string) finalRow {
	splitName := strings.Split(arrRow[0], " ")
	nameToWrite := processName(splitName)
	email := "null"
	domain := "null"

	if len(arrRow[1]) > 0 {
		email = arrRow[1]
		domain = strings.Split(arrRow[1], "@")[1]
	}

	return finalRow{
		firstName: nameToWrite[0],
		lastName:  nameToWrite[1],
		userName:  nameToWrite[2],
		email:     email,
		domain:    domain,
		date:      arrRow[2],
	}

}
