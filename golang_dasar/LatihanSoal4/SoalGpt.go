package main

/*
1. Variabel, Konstanta, dan Tipe Data
Buat sebuah program yang meminta input nama, umur, dan status pekerjaan
dari pengguna, lalu tampilkan pesan yang memuat informasi tersebut menggunakan
konstanta untuk format outputnya.
*/
import "fmt"

type data struct {
	Nama, StatusPekerjaan string
	Umur                  int
}

func main() {
	var nama, statusPekerjaan string
	var umur int
	fmt.Print("Masukkan Nama anda: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan umur: ")
	fmt.Scanln(&umur)
	fmt.Print("Masukkan Status Pekerjaan: ")
	fmt.Scanln(&statusPekerjaan)

	const formatOutput = "Nama %s, Umur %d , pekerjaan %s"
	fmt.Printf(formatOutput, nama, umur, statusPekerjaan)
}
