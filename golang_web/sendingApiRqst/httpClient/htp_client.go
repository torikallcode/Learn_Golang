package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// url endpoint
	url := "https://jsonplaceholder.typicode.com/posts"

	// Data JSON yang akan dikirim
	data := map[string]interface{}{
		"title":  "Belajar http.Client",
		"body":   "Penggunaan dasar",
		"userId": 1,
	}

	// Encode data ke format JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Buat request POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	// Menambahkan Header
	req.Header.Set("Content-Type", "application/json")

	//Mengonfigurasi http.CLient dengan timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Mengirim request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Baca respons
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Body:", string(body))

}
