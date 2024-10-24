package main

import "fmt"

/*
Golang , selain merupakan bahasa pemrograman, golang juga menyediakan standard library
(package bawaan) tanpa harus menggunakan package dari luar buatan orang lain,
pada materi ini , kita akan bahas lebih detail package package yang terdapat di golang
yang sering di gunakan saat membuat aplikasi
*/

// Package fmt
/*
Package pertama yang sering digunakan yaitu package "fmt" dengan menggunakan function Print,
selain println masih banyak function yang terdapat di package fmt. contohnya yaitu println , printf , scanln
*/

func main() {
	// firtsName := "M. Torikal"
	// lastName := "Akbar"

	// fmt.Println("Hello! '", firtsName, lastName, "'") //menggunakan Println
	// fmt.Printf("Hello! '%s %s'", firtsName, lastName) //menggunakan  Printf

	var name string
	var age int
	fmt.Print("Masukkan nama anda: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan umur anda: ")
	fmt.Scanln(&age)
	fmt.Printf("Hallo %s anda berumur %d tahun", name, age)
}
