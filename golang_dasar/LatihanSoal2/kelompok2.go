package main

import "fmt"

func main() {

	/*
		Soal 1
		Buatlah program yang melakukan komparasi antara
		dua angka dan menampilkan hasilnya (benar atau salah).
	*/
	// a := 2
	// b := 3
	// fmt.Println(a < b)
	// fmt.Println(a > b)

	/*
		Soal 2
		Buatlah program yang menggunakan operasi boolean untuk
		memeriksa dua kondisi. Misalnya, periksa apakah sebuah angka
		lebih besar dari 10 dan kurang dari 20.
	*/
	// c := 13
	// fmt.Println(c > 10 && c < 20)

	/*
		Soal 3
		Buatlah program yang meminta input dari pengguna berupa dua angka.
		Gunakan operasi boolean dan komparasi untuk memeriksa apakah angka pertama
		lebih besar dari angka kedua dan apakah salah satu dari angka tersebut genap.
		Tampilkan hasilnya.
	*/
	fmt.Println("Masukkan angka pertama:  ")
	var inputPertama int
	fmt.Scanln(&inputPertama)
	fmt.Println("Masukkan angka kedua:  ")
	var inputKedua int
	fmt.Scanln(&inputKedua)

	komparasi := inputPertama > inputKedua
	angkaGenap := (inputPertama%2 == 0) || (inputKedua%2 == 0)

	fmt.Println(komparasi)
	fmt.Println(angkaGenap)
}
