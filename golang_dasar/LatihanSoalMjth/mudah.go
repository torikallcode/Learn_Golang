package main

import "fmt"

// Type Deklarasi
type Age int

// Function & Parameter
// func square(a int) int {
// 	return a * a
// }

// function return value
// func isEven(a int) bool {
// 	if a%2 == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// return Multiple Value
// func swap(a, b int) (int, int) {
// 	return b, a
// }

// Named Return value
// func sumAndProduct(a, b int) (penjumlahan, perkalian int) {
// 	penjumlahan = a + b
// 	perkalian = a * b
// 	return penjumlahan, perkalian
// }

//Variadic Function
// func sumAll(a ...int) int {
// 	total := 0
// 	for _, h := range a {
// 		total += h
// 	}
// 	return total
// }

// Function as value
// func isPrime(a int) bool {
// 	if a <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= a; i++ {
// 		if a%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// Function as parameter
// func penjumlahan(a, b int) int {
// 	return a + b
// }
// func perkalian(a, b int) int {
// 	return a * b
// }
// func operasi(f func(int, int) int, a, b int) int {
// 	return f(a, b)
// }

// Recursive Function
// func faktorial(a int) int {
// 	if a == 1 {
// 		return 1
// 	} else {
// 		return a * faktorial(a-1)
// 	}
// }

// Closure
// func accumulator() func() int {
// 	jumlah := 0
// 	return func() int {
// 		jumlah++
// 		return jumlah
// 	}
// }

// Defer
// func openFile() {
// 	defer fmt.Println("File akan ditutup")
// 	fmt.Println("Pesan telah dibuka")
// }

// Panic & Recover
// func selamatkanDiri() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recover dari panic, pesan: ", r)
// 	}
// }

// func divide(a, b float64) float64 {
// 	defer selamatkanDiri()
// 	if b == 0 {
// 		panic("Angka tidak valid")
// 	}
// 	return a / b
// }

// Struct
// type data struct {
// 	Name, Grade string
// 	Age         int
// }

// Interface
// type Shape interface {
// 	luas() float64
// 	keliling() float64
// }

// type Circle struct {
// 	diameter float64
// }

// type Rectangle struct {
// 	height, width float64
// }

// func (c Circle) luas() float64 {
// 	const pi = 3.14
// 	return pi * (c.diameter / 2)
// }

// func (c Circle) keliling() float64 {
// 	const pi = 3.14
// 	return pi * c.diameter
// }

// func (r Rectangle) luas() float64 {
// 	return r.height * r.width
// }

// func (r Rectangle) keliling() float64 {
// 	return 2 * (r.height + r.width)
// }

// func printShapeDetail(s Shape) {
// 	fmt.Printf("Area: %.1f, Perimeter: %.1f\n", s.luas(), s.keliling())
// }

// Nill dan Type assertions
// func printType(i interface{}) {
// 	str, ok := i.(string)
// 	if ok {
// 		fmt.Println("Ini adalah string:", str)
// 		return
// 	}

// 	angka, ok := i.(int)
// 	if ok {
// 		fmt.Println("Ini adalah integer:", angka)
// 		return
// 	}

// 	decimal, ok := i.(float64)
// 	if ok {
// 		fmt.Println("Ini adalah FLoat:", decimal)
// 		return
// 	}

// 	fmt.Println("Ini bukan string , integer , ataupun float")
// }

// Poiter Function
// func ubahNilai(x *int) {
// 	*x = *x + 2
// }

// func kaliLima(x *int) {
// 	*x = *x * 5
// }

// func kuadrat(x *int) {
// 	*x = *x * *x
// }

// func rumusTekaTeki(x *int) {
// 	kaliLima(x)
// 	ubahNilai(x)
// 	kuadrat(x)
// }

// Error
// func validateAge(age int) error {
// 	if age < 18 {
// 		return errors.New("Anda masih dibawah umur")
// 	}
// 	return nil
// }

// func tentu(x bool) error {
// 	if !x {
// 		return errors.New("anda salah")
// 	}
// 	return nil
// }

// func ganjil(x int) error {
// 	if x%2 == 0 {
// 		return errors.New("angka yang anda masukkan genap , harap masukkan angka ganjil")
// 	}
// 	return nil
// }

func main() {

	angka := []int{100, 100, 100, 100, 100}
	total := 0
	rata_rata := 0
	for _, n := range angka {
		total += n
		rata_rata += n / len(angka)
	}

	fmt.Println(total)
	fmt.Println(rata_rata)
	fmt.Println(len(angka))

	buku := map[int]int{
		1: 1000,
		2: 2000,
		3: 3000,
	}

	total_buku := 0
	for index, n := range buku {
		fmt.Println("buku ke:", index, "adalah", n)
		total_buku += n
	}
	fmt.Println("Harga semua buku:", total_buku)
	/*
		umurAnda := 19
		if err := validateAge(umurAnda); err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			fmt.Println("Anda sudah legal , umur anda", umurAnda)
		}

		jawaban := false
		if err := tentu(jawaban); err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			fmt.Println("Jawaban yang anda berikan benar")
		}

		angkaAnda := 12
		if err := ganjil(angkaAnda); err != nil {
			fmt.Println("Error: ", err.Error())
		} else {
			fmt.Println("Angka yang anda masukkan adalah", angkaAnda)
		}
	*/
	/*
		// variable
		angka := 1
		benar := true
		text := "Ini Tulisan"

		fmt.Println(angka)
		fmt.Println(benar)
		fmt.Println(text)
	*/
	/*
		// Const
		const pi = 3.14
		diameter := 10
		keliling := 2 * pi * (float64(diameter) / 2)
		fmt.Println(keliling)
	*/
	/*
		// Conversi
			S_angka := "123"
			i_angka, _ := strconv.Atoi(S_angka)
			result := i_angka + 10
			fmt.Println(result)
	*/
	/*
		// Type Deklarasi
		var umur Age = 20
		fmt.Println(umur)
	*/
	/*
		// Operator
		a := 2
		b := 3
		penjumlahan := a + b
		pengurangan := a - b
		perkalian := a * b
		pembagian := a / b
		fmt.Println(penjumlahan)
		fmt.Println(pengurangan)
		fmt.Println(perkalian)
		fmt.Println(pembagian)
	*/
	/*
		// Perbandingan
		a := 2
		b := 3
		result := a > b
		fmt.Println(result)
	*/
	/*
		// Array
		angka := [5]int{1, 2, 3, 4, 5}
		fmt.Println(angka)
	*/
	/*
		//Slice
		hari := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
		hari_libur := append(hari, "Sabtu Libur", "Minggu Libur")
		fmt.Println(hari_libur)
	*/
	/*
		//Map
		buah := map[string]int{
			"Mangga": 10000,
			"Apel":   13000,
			"Pisang": 15000,
		}
		daftar_buah := buah
		fmt.Println(daftar_buah)
		harga_mangga := buah["Mangga"]
		fmt.Println("Harga dari buah Mangga adalah", harga_mangga)
	*/

	/*
		// If ELse
		score := 100
		if score > 70 {
			fmt.Println("Lulus")
		} else {
			fmt.Println("Tidak Lulus")
		}
	*/
	/*
			// Switch
				hari := 9
		switch hari {
		case 1:
			fmt.Println("Senin")
		case 2:
			fmt.Println("Selasa")
		case 3:
			fmt.Println("Rabu")
		case 4:
			fmt.Println("Kamis")
		case 5:
			fmt.Println("Jumat")
		case 6:
			fmt.Println("Sabtu")
		case 7:
			fmt.Println("Minggu")
		default:
			fmt.Println("Masukkan angka yang valid")
		}
	*/
	/*
		// For loop
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
	*/
	/*
		// Break & Continue
		for i := 1; i <= 10; i++ {
			if i == 5 {
				continue
			}
			if i == 8 {
				break
			}
			fmt.Println(i)
		}
	*/
	/*
		// Function & parameter
		hasil_kuadrat := square(2)
		fmt.Println(hasil_kuadrat)
	*/
	/*
		// Function return value
		angka := isEven(3)
		fmt.Println(angka)
	*/
	/*
		// Function Multiple return value
		angka_pertama, angka_kedua := swap(2, 3)
		fmt.Println(angka_pertama, angka_kedua)
	*/
	/*
		//Named Return Value
		tambah, kali := sumAndProduct(7, 9)
		fmt.Println(tambah, kali)
	*/
	/*
		// Variadic function
		kumpulan_angka := []int{1, 2, 3, 4, 5}
		total_semua := sumAll(kumpulan_angka...)
		fmt.Println(total_semua)
	*/
	/*
		// Function as value
		// bilangan_prima := isPrime
		// fmt.Println(bilangan_prima(1))
	*/
	/*
		// Function as parameter
		fmt.Println(operasi(penjumlahan, 2, 3))
		fmt.Println(operasi(perkalian, 2, 3))
	*/
	/*
		// Anonyomous function
		ganjil := func(a int) {
			for i := 1; i <= a; i++ {
				if i%2 != 0 {
					fmt.Println(i)
				}
			}
		}
		ganjil(9)
	*/
	/*
		// Faktorial
		fmt.Println(faktorial(3))
		hitung := accumulator()
		fmt.Println(hitung())
	*/
	/*
		// Defer
		openFile()
	*/
	/*
		// Panic & recover
			fmt.Println("Hasil bagi", divide(3, 0))
	*/
	/*
		// Struct
			ceo1 := data{
			Name:  "Torikal",
			Age:   20,
			Grade: "A",
		}
		fmt.Printf("Nama %s umur %d dengan grade %s", ceo1.Name, ceo1.Age, ceo1.Grade)
	*/
	/*
		// Pointer Function
		angka := 10
		fmt.Println("Nilai sebelum diubah", angka)

		kaliLima(&angka)
		ubahNilai(&angka)
		kuadrat(&angka)
		rumusTekaTeki(&angka)
		fmt.Println("NIlai setelah diubah", angka)
	*/

	/*
		// Pointer
		angka := 10
		fmt.Println("Angka awal", angka)

		var ptr *int = &angka
		fmt.Println("Alamat Pointer", ptr)
		fmt.Println("Nilai yang ditunjuk pointer", *ptr)

		*ptr = 20
		fmt.Println("Nilai setelah diubah melalui pointer", angka)
	*/
	/*
		// Nil & type assertion
		printType("Hallo")
		printType(1)
		printType(2.3)
		printType(true)
	*/
	/*
		// interface
		circle := Circle{10}
		recetangle := Rectangle{4, 5}
		printShapeDetail(circle)
		printShapeDetail(recetangle)
	*/
}
