package main

import (
	"fmt"
	"slices"
)

// 6. Perbandingan
func compareStrings(a, b string) string {
	if len(a) > len(b) {
		return "String satu lebih panjang"
	} else if len(a) < len(b) {
		return "String dua lebih panjang"
	} else {
		return "Kedua string memiliki panjang yang sama"
	}
}

// 9. If
func ticket(age int) {
	if age < 5 {
		fmt.Println("Tiket gratis")
	} else if age < 17 {
		fmt.Println("Harga tiket adalah $10")
	} else {
		fmt.Println("Harga tiket adalah $20")
	}
}

// 10. switch
func montf(m int) {
	switch m {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("Maret")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("Mei")
	case 6:
		fmt.Println("Juni")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("Agustus")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("Oktober")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("Desember")
	default:
		fmt.Println("Masukkan angka yang valid")
	}
}

func faktorial(n float64) float64 {
	if n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func permutation(n, r float64) float64 {
	result := faktorial(n) / faktorial(n-r)
	return result
}

func combination(n, r float64) float64 {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func main() {
	// 1.Variables, Integers, Booleans, Strings
	// var name string
	// var age int
	// fmt.Print("Input your name: ")
	// fmt.Scanln(&name)
	// fmt.Print("Input your age: ")
	// fmt.Scanln(&age)

	// fmt.Printf("Halo ,%s! your age is %d tahun", name, age)

	// 2.Constants
	// const conMil = 1.60934
	// var mil float64
	// fmt.Print("Masukkan jarak dalam satuan mil: ")
	// fmt.Scanln(&mil)
	// convert := mil * conMil
	// fmt.Printf("%2.f mil = %f km", mil, convert)

	// 3.Konversi Tipe Data
	// var c float64
	// fmt.Print("Masukkan suhu dalam satuan celcius: ")
	// fmt.Scanln(&c)

	// fahrenheit := (c * 9 / 5) + 32
	// kelvin := c + 273.15

	// fmt.Printf("%2.f celcius = %2.f fahrenheit & %2.f kelvin", c, fahrenheit, kelvin)

	// 4.Type Deklarasi
	// student := []Student{
	// 	{"Muhammad", 20},
	// 	{"Torikal", 20},
	// 	{"Akbar", 20},
	// }

	// for _, a := range student {
	// 	fmt.Println(a.Name, a.Age)
	// }

	// 5. Operasi Matematika
	// angka := []int{1, 2, 3, 4, 5}
	// rata_rata := 0
	// for _, n := range angka {
	// 	rata_rata += n
	// }
	// fmt.Print(rata_rata)

	// 6. Perbandingan
	// string1 := "hallooooo"
	// string2 := "halo"
	// fmt.Println(len(string1))
	// fmt.Println(len(string2))
	// fmt.Println(compareStrings(string1, string2))

	// 7. Array
	// var angka [5]int
	// for i := 0; i < len(angka); i++ {
	// 	fmt.Printf("Masukkan angka ke-%d: ", i+1)
	// 	fmt.Scanln(&angka[i])
	// }
	// sort.Ints(angka[:])
	// fmt.Println(angka)

	// 8. Slice
	// a := []int{1, 2, 3}
	// b := []int{1, 4, 5}
	// s := append(a, b...)
	// s = slices.DeleteFunc(s, func(n int) bool {
	// 	for i := 0; i < len(s); i++ {
	// 		if s {

	// 		}
	// 	}
	// })
	// slices.Compact(s)
	// fmt.Print(s)

	// 9. if
	// ticket(20)

	// 10. switch
	// var mont int
	// fmt.Print("Masukkan angka 1-12: ")
	// fmt.Scanln(&mont)
	// montf(mont)

	// 11.

	// 12. break & continue
	// for i := 0; i <= 20; i++ {
	// 	if i%2 != 0 {
	// 		fmt.Println(i)
	// 		if i == 13 {
	// 			fmt.Println("Loop berhenti di angka", i)
	// 			break
	// 		}
	// 	}
	// }

	// 13.function & parameter
	// fmt.Println(permutation(9, 8))
	// fmt.Println(combination(9, 8))

	angka := []int{1, 2, 3}
	fmt.Println(slices.Max(angka))
	fmt.Println(slices.Min(angka))

}
