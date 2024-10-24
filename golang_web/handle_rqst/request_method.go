package main

import (
	"fmt"
	"net/http"
)

func handleFunction_requestMethod(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello world")
	case "/profil":
		fmt.Fprintf(w, "Hello Akbar")
	default:
		fmt.Fprint(w, "Error page")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func main() {
	http.HandleFunc("/", handleFunction_requestMethod)
	http.ListenAndServe("", nil)
}
