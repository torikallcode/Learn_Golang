package main

import "fmt"

/*
	Soal 1
	Buatlah fungsi sederhana yang menerima dua parameter
	berupa angka dan mengembalikan hasil penjumlahannya.
*/
func penjumlahan(a, b int) int {
	return a + b
}

/*
	Soal 2
	Buatlah fungsi yang menerima parameter berupa array
	dan mengembalikan jumlah dari semua elemen dalam array tersebut.
*/
func angka(a []int) {
	for _, value := range a {
		fmt.Println(value)
	}
}

/*
	Soal 3
	Buatlah fungsi yang menerima fungsi lain sebagai parameter
	dan sebuah angka. Fungsi ini harus memanggil fungsi yang diterima
	sebagai parameter untuk memproses angka tersebut, lalu mengembalikan hasilnya.
	Misalnya, fungsi utama menerima fungsi yang melakukan perkalian dengan 2,
	dan angka yang akan dikalikan.
*/
type Perkalian func(int) int

func rumus(angkaS int, perkalian Perkalian) int {
	return perkalian(angkaS)
}

func perkalian(angkaD int) int {
	return angkaD * 2
}

func main() {

	// Soal 1
	fmt.Println(penjumlahan(9, 10))

	// soal 2
	number := []int{1, 2, 3, 4, 5}
	angka(number)

	// soal 3
	hasil := rumus(9, perkalian)
	fmt.Println(hasil)

}
