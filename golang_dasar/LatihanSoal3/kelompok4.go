package main

import "fmt"

// Soal 1
func panjang(text string) int {
	return len(text)
}

// Soal 2
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Soal 3
func cariPangkat(angka, pangkat int) int {
	result := 1
	for i := 0; i < pangkat; i++ {
		result *= angka
	}
	return result

}

func main() {
	/*
		Soal 1
		Buatlah sebuah fungsi yang menerima sebuah string
		sebagai parameter dan mengembalikan panjang string tersebut.
	*/
	// fmt.Println(panjang("Hallo"))

	/*
		Soal 2
		Buatlah fungsi yang menerima dua parameter berupa angka.
		Fungsi tersebut harus mengembalikan nilai maksimal dari kedua angka tersebut.
	*/
	// nilaiMax := max(2, 3)
	// fmt.Println("Nilai maksimal adalah", nilaiMax)

	/*
		Soal 3
		Buatlah dua fungsi: fungsi pertama untuk menghitung pangkat dari sebuah angka,
		dan fungsi kedua untuk menghitung hasil akar kuadrat dari sebuah angka.
		Buatlah fungsi ketiga yang menerima salah satu dari kedua fungsi ini
		sebagai parameter dan sebuah angka. Fungsi ketiga ini harus mengembalikan
		hasil operasi dari fungsi yang diterima sebagai parameter terhadap angka yang diberikan.
	*/

	fmt.Println(cariPangkat(2, 3))
}
