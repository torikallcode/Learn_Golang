package main

import (
	"fmt"
	"strings"
)

// Type deklarasi
// type student struct {
// 	Name string
// 	Age  int
// }

// Perbandingan
// func compareStrings(s1, s2 string) string {
// 	if len(s1) > len(s2) {
// 		return "String 1 lebih panjang"
// 	} else if len(s1) < len(s2) {
// 		return "String 2 lebih panjang"
// 	} else {
// 		return "Kedua String memiliki panjang yang sama"
// 	}
// }

// FUnction & Parameter
// func faktorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	return result
// }

// func fk(n int) int {
// 	if n == 1 {
// 		return 1
// 	} else {
// 		return n * fk(n-1)
// 	}
// }

// func permutation(n, r int) int {
// 	return faktorial(n) / faktorial(n-r)
// }

// func combination(n, r int) int {
// 	return faktorial(n) / (faktorial(r) * faktorial(n-r))
// }

// function return value

func uppercase(str string) string {
	return strings.ToUpper(str)
}

func main() {

	var inputStr string
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&inputStr)

	fmt.Println(uppercase(inputStr))

	/*
		// Variable
		var nama string
		var umur int
		fmt.Print("Masukkan nama anda: ")
		fmt.Scanln(&nama)
		fmt.Print("Masukkan umur anda: ")
		fmt.Scanln(&umur)
		fmt.Printf("Halo, %s! Umur kamu %d tahun", nama, umur)
	*/
	/*
		// Constant
		var jarak float64
		const km = 1.60934
		fmt.Print("Masukkan jarak dalam bentuk mil: ")
		fmt.Scanln(&jarak)
		konversi := jarak * km
		fmt.Printf("%v mil = %vkm", jarak, konversi)
	*/
	/*
		// Konversi tipe data
		var celcius int
		fmt.Print("Masukkan suhu dalam satuan celcius: ")
		fmt.Scanln(&celcius)
		konversi_celcius := float64(celcius)
		var fahrenheit float64 = (konversi_celcius * 9 / 5) + 32
		var kelvin float64 = konversi_celcius + 273.15
		fmt.Printf("%v Celcius = %v fahrenheit\n", celcius, fahrenheit)
		fmt.Printf("%v Celcius = %v kelvin", celcius, kelvin)
	*/
	/*
		// Type deklarasi
		student := []student{
			{"Muhammad", 20},
			{"Torikal", 20},
			{"Akbar", 20},
		}

		fmt.Println(student)
		for _, nama := range student {
			fmt.Printf("Nama: %v , Umur: %v\n", nama.Name, nama.Age)
		}
	*/
	/*
		// Operasi matematika
		var total_nilai int
		fmt.Print("Masukkan jumlah angka: ")
		fmt.Scanln(&total_nilai)
		nilai := make([]int, total_nilai)
		rata_rata := 0
		for i := 0; i < total_nilai; i++ {
			fmt.Printf("Masukkan angka ke-%d: ", i+1)
			fmt.Scanln(&nilai[i])
			rata_rata += nilai[i]
		}
		average := float64(rata_rata) / float64(total_nilai)
		fmt.Printf("Rata rata: %.2f\n", average)
	*/
	/*
		// Perbandingan
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan kata pertama: ")
		kata1, _ := reader.ReadString('\n')
		kata1 = strings.TrimSpace(kata1) // Menghapus karakter newline di akhir input
		fmt.Print("Masukkan kata kedua: ")
		kata2, _ := reader.ReadString('\n')
		kata2 = strings.TrimSpace(kata2) // Menghapus karakter newline di akhir input
		fmt.Println(compareStrings(kata1, kata2))
		fmt.Println(compareStrings("Halo bro", "Dunia"))
		fmt.Println(compareStrings("Selamat", "Pagi bro"))
		fmt.Println(compareStrings("Go", "Go"))
	*/
	/*
		// Array
		var inputAngka [5]int
		for i := 0; i < len(inputAngka); i++ {
			fmt.Printf("Masukkan angka ke-%d: ", i+1)
			fmt.Scanln(&inputAngka[i])
		}
		sort.Ints(inputAngka[:])
		fmt.Println("Array setelah diurutkan", inputAngka)
	*/
	/*
		// Slice
	*/
	/*
		// Map
		var nilai_mhs = map[string]int{
			"Muhammad": 100,
			"Torikal":  100,
			"Akbar":    100,
		}
		fmt.Println("Daftar siswa dan nilainya: ")
		for name, score := range nilai_mhs {
			fmt.Println(name, score)
		}

		nilai_mhs["Evand"] = 100
		fmt.Println(nilai_mhs)
		delete(nilai_mhs, "Evand")
		fmt.Println(nilai_mhs)
	*/
	/*
		// If Statement
		var usia int
		fmt.Print("Masukka usia mu: ")
		fmt.Scanln(&usia)

		if usia < 5 {
			fmt.Println("Tiket Gratis")
		} else if usia <= 17 {
			fmt.Println("Tiket seharga 10$")
		} else {
			fmt.Println("Tiket seharga 20$")
		}
	*/
	/*
		// Swith statement
		var angka int
		fmt.Print("Masukkan antar angka 1 - 12: ")
		fmt.Scanln(&angka)

		switch angka {
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
			fmt.Println("Angka tidak Valid")
		}
	*/
	/*
		// Break & Continue
		for i := 0; i <= 20; i++ {
			if i%2 == 0 {
				continue
			}
			if i == 13 {
				fmt.Println("Loop dihentikan di angka 13")
				break
			}
			fmt.Println(i)
		}
	*/
	/*
		// Function parameter
			fmt.Println(permutation(6, 2))
		fmt.Println(combination(6, 2))
	*/

}
