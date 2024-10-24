package main

import (
	"flag"
	"fmt"
)

func main() {
	// username := flag.String("username", "root", "database username")
	// password := flag.String("password", "root", "database password")
	// host := flag.String("host", "localhost", "database host")
	// port := flag.Int("port", 0, "55005")

	// flag.Parse()

	// fmt.Println("Username", *username)
	// fmt.Println(*password)
	// fmt.Println(*host)
	// fmt.Println(*port)

	name := flag.String("name", "anonim", "Masukkan nama anda")
	age := flag.Int("age", 0, "Masukkan nama anda")
	married := flag.Bool("married", false, "Apakah anda sudah menikah")

	flag.Parse()
	fmt.Printf("Nama:%s , Umur:%d , Menikah:%t", *name, *age, *married)
}

/*
// Package flag
package flag berisikan fungsionalitas untuk memparsing command line argument
contoh
username := flag.String("username", "root", "database username")
"username" = nama flagnya
"root" = default value
"database username" = deskripsi

flag.Parse() = untuk melakukan parsing

balikkannya adalah pointer , jadi jika di panggil maka seperti ini( menggunakan asterisk)
fmt.Println("Username",*username)
// menjalankan aplikasi
misal hanya flag.go maka akan menampilkan default valuenya
output = Username root
jika ingin mengganti maka bisa mengguakan
flag.go -username=Akbar
jika lebih dari satu kata maka menggunakan string
flag.go -username="Torikal Akbar"
*/
