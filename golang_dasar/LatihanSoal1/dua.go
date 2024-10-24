// SOAL 2: Waktu 15 Menit, 25 poin
// Keliling lingkaran didapatkan dengan rumus: K = PI * diameter, dimana diameter adalah 2 kali jari-jari
// Luas lingkaran didapatkan dengan rumus: L = PI * jari-jari * jari-jari
// Dimana PI = 3.14 atau 22/7
// Buatlah sebuah program yang menerima input jari-jari dari user kemudian mencetak nilai keliling dan luasnya ke layar

package main

import "fmt"

// menggunakan function
func keliling(jari_jari float64) float64 {
	pi := 3.14
	hasil := pi * jari_jari * 2
	return hasil
}

func luas(jari_jari float64) float64 {
	pi := 3.14
	return pi * jari_jari * jari_jari
}


func main() {
	// langsung dalam function main
	// jari_jari := 23
	// keliling := 3.14 * ((float64(jari_jari)) * 2)
	// luas := 3.14 * (float64(jari_jari)) * (float64(jari_jari))

	// fmt.Println("Keliling lingkaran = ",keliling)
	// fmt.Println("Luas lingkaran = ",luas)
	fmt.Println(keliling(23))
	fmt.Println(luas(23))
}