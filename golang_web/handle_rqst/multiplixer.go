package main

import (
	"fmt"
	"net/http"
	"time"
)

func htmlVsPlain3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func timeout3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("timeout")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not time")
}

func helloWorldNinjaMode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloWorldNinjaMode")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello World</h1>")
}

func main() {
	http.HandleFunc("/", timeout3)
	// http.HandleFunc("/timeout", timeout3)

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxNinjaMode http.ServeMux
	server.Handler = &muxNinjaMode
	muxNinjaMode.HandleFunc("/", helloWorldNinjaMode)

	server.ListenAndServe()

}
