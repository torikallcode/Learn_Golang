package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "M. Torikal akbar"
	encode := base64.StdEncoding.EncodeToString([]byte(value)) //membuat encoding menggunakan base64
	fmt.Println(encode)

	decode, err := base64.StdEncoding.DecodeString(encode) //melakukan decode pada encode
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(string(decode))
	}
}
