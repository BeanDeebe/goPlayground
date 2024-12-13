package main

import (
	"encoding/csv"
	"os"
)

func main() {

	fname, err := os.Open("data/data.csv")
	

	if err != nil {
		panic(err)
	}

	defer fname.Close()

	r := csv.NewReader(fn
}
