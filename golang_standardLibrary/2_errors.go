package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("Validation error")
	notFoundError   = errors.New("NotFound Error")
)

func getById(n string) error {
	if n == "" {
		return validationError
	}

	if n != "Akbar" {
		return notFoundError
	}
	return nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagi tidak bisa berupa 0")
	}
	return a / b, nil
}

func multiplication(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Perkalian tidak bisa menggunakan 0")
	}
	return a * b, nil
}

func main() {
	// var id string
	// fmt.Print("Masukkan ID anda: ")
	// fmt.Scanln(&id)

	// err := getById(id)
	// if err != nil {
	// 	if errors.Is(err, validationError) {
	// 		fmt.Println("Validation Error")
	// 	} else if errors.Is(err, notFoundError) {
	// 		fmt.Println("NotFound Error")
	// 	} else {
	// 		fmt.Println("Unknown Error")
	// 	}
	// } else {
	// 	fmt.Printf("ID anda adalah: %s", id)
	// }

	// var a, b int
	// fmt.Print("Masukkan angka pertama: ")
	// fmt.Scanln(&a)
	// fmt.Print("Masukkan angka kedua: ")
	// fmt.Scanln(&b)

	// result, err := divide(a, b)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// } else {
	// 	fmt.Println("Hasil:", result)
	// }

	var a, b int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	result, err := multiplication(a, b)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Hasil perkalian", result)
	}

}

/*
Di materi golang dasar kita sudah membahasa interface error yang merupakan representasi
dari error di golang, dan membuat error menggunakan fucntion error.New(),
Sebenarnya masih banyak yang bisa kita lakukan menggunakan package errors, contohnya
ketika kita ingin beberapa value error yang berbeda
kebiasaan programer golang yaitu dalam sebuah variable atau constant =
var (
	validationError = errors.New("Validation error")
	notFoundError   = errors.New("Nots Found error")
)
jadinya jika kita butuh error dalam bentuk validation error misalnya , kita tinggal panggil
variable tersebut.
Misal kita membuat error sendiri, lalu kita ingin mengecek jenis errornya , kita bisa
menggunakan error.Is() untuk mengecek type errornya =
if errors.Is(err, validationError) {
			fmt.Println("Validation Error")
		}
di cek apakah errornya adalah validation error , jika iya maka print "validation error"

*/

// SOAL
/*
1. Buatlah sebuah program operasi pembagian yang dimana ketika pembaginya adalah 0
maka berikan pesan error
2. Buatlah sebuah program yang menerima input id dari pengguna , jika id adalah
string kosong maka berikan pesan error validation error , jika id tidak sama dengan
Akbar maka berikan error notfound error
*/
