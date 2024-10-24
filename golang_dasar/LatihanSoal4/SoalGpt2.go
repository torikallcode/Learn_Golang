package main

import (
	"fmt"
)

/*
Bagus sekali! Berikut adalah beberapa latihan untuk setiap konsep fungsi di Go yang sudah kamu pelajari:

### 1. **Function dengan Parameter dan Return Value**
  - Buat fungsi bernama `add` yang menerima dua parameter integer dan mengembalikan jumlah keduanya.
  - Buat fungsi bernama `subtract` yang menerima dua parameter integer dan mengembalikan selisih keduanya.

### 2. **Returning Multiple Values**
  - Buat fungsi bernama `divideAndRemainder` yang menerima dua parameter integer
    dan mengembalikan hasil bagi dan sisa pembagian keduanya.

### 3. **Named Return Value**
  - Buat fungsi bernama `rectangleProperties` yang menerima panjang dan lebar dan
    mengembalikan luas dan kelilingnya menggunakan named return values.

### 4. **Variadic Function**
  - Buat fungsi bernama `sumAll` yang menerima parameter variadik integer dan
    mengembalikan jumlah total semua angkanya.

### 5. **Function as Value**
  - Buat sebuah fungsi bernama `multiply` dan simpan ke dalam variabel.
    Kemudian, panggil fungsi tersebut menggunakan variabel tersebut.

### 6. **Function as Parameter**
  - Buat fungsi bernama `applyOperation` yang menerima dua integer dan
    sebuah fungsi sebagai parameter. `applyOperation` akan menerapkan fungsi tersebut ke dua angka tersebut.

### 7. **Anonymous Function**
  - Buat sebuah fungsi anonim yang mengalikan dua angka dan panggil
    fungsi tersebut segera setelah definisinya.

### 8. **Recursive Function**
  - Buat fungsi bernama `factorial` yang menerima integer dan
    mengembalikan nilai faktorialnya menggunakan rekursi.

### 9. **Closure**
  - Buat fungsi bernama `counter` yang mengembalikan fungsi lain.
    Fungsi yang dikembalikan akan menambah nilai counter setiap kali dipanggil.

### 10. **Defer**
  - Buat fungsi bernama `logMessage` yang menggunakan kata kunci `defer`
    untuk mencetak "Selesai!" setelah mencetak pesan yang diberikan sebagai argumen.

### 11. **Panic dan Recover**
  - Buat fungsi bernama `divide` yang mengembalikan hasil pembagian dua angka.
    Jika angka pembaginya adalah nol, gunakan `panic`, dan gunakan `recover` di fungsi yang memanggilnya untuk menangani panic tersebut.
*/
func selamatkanDiri() {
	if r := recover(); r != nil {
		fmt.Println("Recover dari panic , pesan:", r)
	}
}

func divide(x, y int) int {
	defer selamatkanDiri()
	if y == 0 {
		panic("Angka yang anda masukkan 0")
	}
	return x / y
}

// SOAL 1
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

// SOAl 2
func divideAndRemainder(a, b int) (int, int) {
	return a / b, a % b
}

// SOAL 3
func rectangleProperties() (panjang, lebar int) {
	panjang = 10
	lebar = 10
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

// SOAL 4
func sumAll(a ...int) float64 {
	total := 0
	for _, i := range a {
		total += i
	}
	return float64(total)
}

// SOAL 5
func multiply(nama string) string {
	return "Hallo " + nama
}

//SOAL6

type Penjumlahan func(int, int) int

func applyOperation(a int, b int, penjumlahan Penjumlahan) {
	fmt.Println("Hasil:", penjumlahan(a, b))
}

func pemjumlahan(a, b int) int {
	return a + b
}

// SOAL 7

// SOAL8
func faktorial(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * faktorial(a-1)
	}
}

// Soal 10
func endApp() {
	fmt.Println("Program Selesai")
}

func sayHello(name string) {
	defer endApp()
	fmt.Println("Hello", name)
}

func main() {

	fmt.Println(divide(2, 0))

	// SOAL 1
	// fmt.Println(add(2, 3))
	// fmt.Println(subtract(4, 4))

	// SOAL 2
	// fmt.Println(divideAndRemainder(3, 2))

	// SOAL 3
	// a, b := rectangleProperties()
	// fmt.Println(a, b)

	// SOAL 4
	// totalAngka := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// fmt.Println(totalAngka)

	// SOAL 5
	// sayHello := multiply
	// fmt.Println(sayHello("Akbar"))

	// SOAL 6
	// applyOperation(2, 3, pemjumlahan)

	// SOAL 7
	// perkalian := func(a, b int) int {
	// 	return a * b
	// }

	// fmt.Println(perkalian(3, 2))

	//SOAL 8
	// fmt.Println(faktorial(1))

	// SOAL 9
	// counter := 0

	// tambah := func() {
	// 	counter++
	// }

	// tambah()
	// fmt.Println(counter)

	// SOAL 10
	// sayHello("Akbar")

}
