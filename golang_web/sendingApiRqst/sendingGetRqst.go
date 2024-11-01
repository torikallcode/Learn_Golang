package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Mengirim permintaan GET
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/2") //url yang akan diakses
	if err != nil {
		log.Fatal(err) //log.Fatal
	}
	defer response.Body.Close()
	/*memastikan bahwa koneksi body respons ditutup setelah selesai dibaca.
	Ini penting untuk menghindari kebocoran memori (memory leak).
	*/

	// Membaca response body
	/*
		body, err := ioutil.ReadAll(response.Body) membaca seluruh isi respons dan
		menyimpannya dalam variabel body.
	*/
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
