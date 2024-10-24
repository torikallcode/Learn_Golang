package main

import "fmt"

func main() {

	// Membuat type data baru atau Alias dari string
	type noKtp string

	// variable ktpSaya bertipe data noKtp alias string
	var ktpSaya noKtp = "111111111"

	// cara konversi tipe data yang sudah digunakan sebelumnya menjadi tipe data noKtp
	var contoh string = "222222222"
	var contohKtp noKtp = noKtp(contoh)

	fmt.Println(ktpSaya)
	fmt.Println(contohKtp)


}