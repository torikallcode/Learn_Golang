package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	jsonData := []byte(`{"title": "react" , "body":"state management" , "userID": 1}`)

	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/users/1", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hasil Put: ", string(body))
}
