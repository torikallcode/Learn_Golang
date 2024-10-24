package main

import (
	"fmt"
	"net/http"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello world")
	case "/ninja":
		fmt.Fprint(w, "Hello Akbar")
	default:
		fmt.Fprint(w, "Page Error")
	}
}

func main() {
	http.HandleFunc("/", handleFunction)
	http.ListenAndServe("", nil)
}
