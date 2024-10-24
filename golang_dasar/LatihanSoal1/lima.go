// Bangun ruang berbentuk silinder dengan jari-jari r, dan tinggi t
// diketahui memiliki volume sebesar πr2t dan luas permukaan sebesar 2πr2 + 2πrt
// Buatlah sebuah program yang menerima input jari-jari dan tinggi silinder, kemudian mencetak volume dan luasnya pada layar.
// Gunakan nilai π = 3.14

package main

import "fmt"

func volumeSilinder(tinggi , jari_jari float64) float64 {
	pi := 3.14
	return pi * (jari_jari * jari_jari) * tinggi
}

func luas_permukaanSilinder(tinggi , jari_jari float64) float64 {
	pi := 3.14
	return (2* pi * (jari_jari * jari_jari)) + (2 * pi * jari_jari * tinggi)
}

func main() {

	// tinggi := 10.0
	// jari_jari := 15.0
	// const pi = 3.14

	// volume := pi * (jari_jari * jari_jari) * tinggi
	// luas_permukaan := (2 * pi *(jari_jari * jari_jari)) + (2 * pi * jari_jari * tinggi)

	// fmt.Println("Volume silinder= ", volume)
	// fmt.Print("Luas permukaan silinder= ", luas_permukaan)

	fmt.Println(volumeSilinder(10.0 , 15.0))
	fmt.Println(luas_permukaanSilinder(10.0 , 15.0))

}