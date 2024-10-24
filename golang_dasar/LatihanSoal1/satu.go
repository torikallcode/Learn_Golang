package main

import "fmt"

// Buatlah variabel yang menyimpan nilai-nilai berikut:

// Nama
// Tahun Lahir
// Sekolah Dasar / MI
// Sekolah Menengah Pertama / MTs
// Sekolah Menengah Atas / SMK / MA
// Cetak nilai seluruh variabel tersebut ke layar

func dataDiri(nama string , thn_lahir int16 , sd string , mts string , sma string) {
	fmt.Println(nama)
	fmt.Println(thn_lahir)
	fmt.Println(sd)
	fmt.Println(mts)
	fmt.Println(sma)
}

func main() {
	// nama := "M. Torikal Akbar"
	// thn_lahir := 2003
	// sd := "SDN 1 Pancor"
	// mts := "MTS N 1 Lotim"
	// SMA := "SMA N 2 Selong"

	// fmt.Println(nama)
	// fmt.Println(thn_lahir)
	// fmt.Println(sd)
	// fmt.Println(mts)
	// fmt.Println(SMA)

	dataDiri("M. Torikal Akbar" , 2003 , "SDN 1 Pancor" , "MTSN 1 Lotim" , "SMAN 2 Selong")
}