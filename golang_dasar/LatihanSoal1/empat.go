// Body Mass Index (indeks massa tubuh), atau BMI dihitung dengan rumus sebagai berikut.

// BMI = bb / tb2

// bb: berat badan dalam satuan kilogram
// tb: tinggi badan dalam satuan meter

// Nilai BMI kemudian diklasifikasikan sebagai berikut:

// Di bawah 17.5, kurus
// 17.5 s.d di bawah 23, normal
// 23 s.d di bawah 28, gemuk
// 28 ke atas, obesitas

// Buatlah sebuah program yang menerima input nama, tinggi badan dalam satuan cm, dan berat badan seseorang dalam satuan kg, kemudian menghitung BMI, menetapkan klasifikasinya, dan mencetak nilai BMI beserta klasifikasinya ke layar.

// Catatan: 1 meter = 100 cm

package main

import "fmt"

func bmi(berat , tinggi float32) (float32, string) {
	tinggiMetter := tinggi /100
	hasilBmi := berat /(tinggiMetter * tinggiMetter)

	status := ""
	if hasilBmi < 17.5 {
		status = "Kurus"
	} else if hasilBmi < 23 {
		status = "Ideal"
	} else if hasilBmi < 28 {
		status = "Gemuk"
	} else {
		status = "Obesitas"
	}
	return hasilBmi , status
}

func main() {

	// berat := 70.0
	// tinggi := 165.0
	// tingiMetter := tinggi / 100
	// bmi := berat / (tingiMetter * tingiMetter)

	// status := ""
	// if bmi < 17.5 {
	// 	status = "Kurus"
	// } else if bmi < 23 {
	// 	status = "Ideal"
	// } else if bmi < 28 {
	// 	status = "Gemuk"
	// } else {
	// 	status = "Obesitas"
	// }

	// fmt.Println(bmi)
	// fmt.Println(status)

	fmt.Println(bmi(50.0 , 165.0))
}