package main

import "fmt"

func main() {
	nama := "Torikal"

	if nama == "Muhammad" {
		fmt.Println("Hallo Muhammad")
	} else if nama == "Torikal" {
		fmt.Println("Hallo Torikal")
	} else if nama == "Akbar" {
		fmt.Println("Hallo Akbar")
	} else {
		fmt.Println("Hallo, boleh kenalan?")
	}

	// panjang := len(nama)

	// if panjang > 5 {
	// 	fmt.Println("Nama terlalu panjang")
	// } else {
	// 	fmt.Println("Nama anda benar")
	// }

	// Menggunakan Sort Statement
	if panjang := len(nama); panjang > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama anda sudah benar")
	}
}

