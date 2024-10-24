package main

import (
	"fmt"
)

func main() {
	nama := "Akbar"

	switch nama {
	case "Muhammad":
		fmt.Println("Hello Muhammad")
	case "Torikal":
		fmt.Println("Hello Torikal")
	case "Akbar" :
		fmt.Println("Hello Akbar")
	default :
		fmt.Println("Hai, Boleh Kenalan?")
	}

	panjang := len(nama)

	switch panjang > 5 {
	case true:
		fmt.Println("Nama Terlalu panjang")
	case false:
		fmt.Println("Nama Sudah benar")
	} 

	nama = "Torikal"
	// Short statement
	switch panjang := len(nama); {
	case panjang > 10 :
		fmt.Println("Nama terlalu Panjang")
	case panjang > 5 :
		fmt.Println("Nama Lumayan Panjang")
	default :
		fmt.Println("Nama sudah benar")
	}

}