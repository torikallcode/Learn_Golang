package main

import "fmt"

func main() {
	// Constant nilainya tidak dapat diubah dan harus langsung di inisialisasikan dan juga tidak terjadi error ketika tidak di panggil
	const firstName = "Muhammad"
	const lastName = "Akbar"

	const (
		midleName = "Torikal"
	)

	fmt.Println(firstName)
	fmt.Println(midleName)
	fmt.Println(lastName)


}