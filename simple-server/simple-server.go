package main

import (
	"io"
	"net/http"
)

func main() {

	hello := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello!")
	}

	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}
