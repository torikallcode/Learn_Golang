package main

import (
	"fmt"
	"os"
)

func main() {

	// args := os.Args
	// for _, arg := range args {
	// 	fmt.Println(arg)
	// }

	// hostName, err := os.Hostname()
	// if err == nil {
	// 	fmt.Println(hostName)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	// file, err := os.Open("data.text")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer file.Close()

	// file, err := os.Create("baru.txt")
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// defer file.Close()

	// file1, err := os.Open("baru.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer file1.Close()

	// err := os.Remove("baru.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	path := os.Getenv("PATH")
	fmt.Println("PATH:", path)

}

/*
// Package os
package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen
(bisa digunakan disemua sistem operasi).
yang pertama yaitu args = yang digunakan untuk mendapatkan argumen ketika kita menjalankan aplikasi
jadi di saat menjalankan aplikasi kita bisa mengirimkan argument dan args ini untuk menangkap argument tersebut

misal go run os.go Muhammad Torikal Akbar
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

output =
// C:\Users\ACERRY~1\AppData\Local\Temp\go-build2551493841\b001\exe\3_os.exe = lokasi file
// lalu parameter yang dikirim
Muhammad
Torikal
Akbar
jika menggunakan String go run os.go "Muhammad Torikal Akbar"
maka akan jadi satu
Muhammad Torikal Akbar
// contoh 2 = untuk memanggil hostname 	hostName, err := os.Hostname()
	if err == nil {
		fmt.Println(hostName)
	} else {
		fmt.Println("Error", err.Error())
	}

outout
Muhammad Torikal Akbar
DARKSIDE
*/
