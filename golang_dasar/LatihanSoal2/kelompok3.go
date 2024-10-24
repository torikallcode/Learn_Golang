package main

import "fmt"

func main() {
	/*
		Soal 1
		Buatlah array dengan 5 elemen bertipe integer dan
		tampilkan semua elemen tersebut.
	*/
	// elemen := []int{1, 2, 3, 4, 5}
	// fmt.Println(elemen)
	/*
		Soal 2
		Buatlah slice dari array yang sudah ada, tambahkan
		beberapa elemen ke dalam slice, dan tampilkan hasil akhirnya.
	*/
	// newElemen := append(elemen, 6, 7, 8, 9, 10)
	// fmt.Println(newElemen)
	/*
		Soal 3
		Buatlah map untuk menyimpan data mahasiswa
		(misalnya nama sebagai key dan nilai sebagai value).
		Tampilkan semua data dalam map tersebut, dan carilah
		nilai rata-rata dari semua nilai mahasiswa yang ada.
	*/
	mahasiswa := map[string]int{
		"Muhammad": 100,
		"Torikal":  100,
		"Akbar":    100,
	}

	for nama, nilai := range mahasiswa {
		fmt.Println(nama, nilai)
	}

	total := 0
	jumlah_mhs := len(mahasiswa)
	for _, nilai := range mahasiswa {
		total += nilai
	}
	rata_rata := total / jumlah_mhs

	fmt.Println(rata_rata)
}
