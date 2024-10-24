package main

import (
	"fmt"
	"net/http"
	"time"
)

func htmlVsPlain2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain2")
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Attemp timeout")
	time.Sleep(30 * time.Second)
	fmt.Fprintf(w, "Did not timeout")
}

func main() {
	http.HandleFunc("/", htmlVsPlain2)
	http.HandleFunc("/timeout", timeout)
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()
}
