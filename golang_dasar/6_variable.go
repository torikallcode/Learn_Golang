package main

import "fmt"

func main() {
	// Cara Pertama
	// Mendefinisikan variabel diikuti dengan nama variabel dan type data
	var name string

	name = "M. Torikal Akbar"
	fmt.Println(name)

	name = "Muhammad Torikal Akbar"
	fmt.Println(name)

	// cara kedua
	// Langsung menaruh nilai pada variable tidak menulis type data 
	// karena golang tau apa type data dari isi variable tersebut
	var alamat = "pancor"
	fmt.Println(alamat)

	// cara ke-3
	// tidak menggunakan "var" dan menggunakan ":=" sebagai deklarasi pertama , 
	// dan jika ingin mengganti isi variable tersebut tidak perlu perlu menggunakan ":" hanya "=" saja
	// Paling sering digunakan
	univ := "Unhaz"
	fmt.Println(univ)

	univ = "Universitas Hamzanwadi"
	fmt.Println(univ)

	// cara ke-4
	// Mendeklarasikan variable secara sekaligu
	// Terjadi error jika tidak memanggil variable 
	var (
		firtsName = "Muhammad"
		midleName = "Torikal"
		lastName = "Akbar"
	)

	fmt.Println(firtsName)
	fmt.Println(midleName)
	fmt.Println(lastName)

}