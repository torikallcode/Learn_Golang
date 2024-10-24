package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		1.Soal pertama
		Buatlah program sederhana yang mendeklarasikan tiga variabel
		dengan tipe data yang berbeda (integer, boolean, dan string).
		Tampilkan nilai variabel tersebut.
	*/
	// integer := 1
	// boolean := true
	// tulisan := "String"
	// fmt.Println(integer)
	// fmt.Println(boolean)
	// fmt.Println(tulisan)

	/*
		2.Soal kedua
		Buatlah program yang melakukan operasi matematika sederhana
		(penjumlahan, pengurangan, perkalian, pembagian) menggunakan d
		ua variabel bertipe integer. Konversikan hasil dari operasi tersebut
		menjadi string dan tampilkan.
	*/
	// a := 2
	// b := 3

	// penjumlahan := a + b
	// pengurangan := a - b
	// perkalian := a * b
	// pembagian := a / b

	// convPenjumlahan := strconv.Itoa(penjumlahan)
	// conPengurangan := strconv.Itoa(pengurangan)
	// convPerkalian := strconv.Itoa(perkalian)
	// convPembagian := strconv.Itoa(pembagian)

	// fmt.Println(convPenjumlahan)
	// fmt.Println(conPengurangan)
	// fmt.Println(convPerkalian)
	// fmt.Println(convPembagian)

	/*
		3.Soal ketiga
		Buatlah program yang menerima input dari pengguna untuk dua angka,
		lakukan operasi matematika pada angka tersebut, dan tampilkan hasilnya
		dengan format string. Pastikan hasil yang ditampilkan tetap benar walaupun
		pengguna memasukkan angka dalam format string.
	*/
	fmt.Println("Masukkan angka pertama:  ")
	var inputPertama string
	fmt.Scanln(&inputPertama)
	angkaPertama, err := strconv.Atoi(inputPertama)
	if err != nil {
		fmt.Println("Harap Masukkan angka")
		return
	}

	fmt.Println("Masukkan angka kedua:  ")
	var inputKedua string
	fmt.Scanln(&inputKedua)
	angkaKedua, err := strconv.Atoi(inputKedua)
	if err != nil {
		fmt.Println("Harap masukkan angka")
	}

	hasil := angkaPertama * angkaKedua
	convPerkalian := strconv.Itoa(hasil)
	fmt.Println(convPerkalian)

}
