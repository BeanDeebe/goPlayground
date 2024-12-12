package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	data, err := os.Open("data/data.csv")
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

		fmt.Println(line)
	}
}
