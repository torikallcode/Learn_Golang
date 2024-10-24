package main

import "fmt"

func main() {

	/*
		Soal 1
		Buatlah array dengan 5 elemen bertipe string yang
		berisi nama hari dalam seminggu, lalu tampilkan nama hari ke-4.
	*/
	// hari := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	// fmt.Println(hari[3])
	/*
		Soal 2
		Buatlah slice dari sebuah array yang berisi 10 angka.
		Hapus 2 angka dari slice tersebut dan tambahkan 3 angka baru, lalu tampilkan hasil akhirnya.
	*/
	// angka := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Sl_angka := angka[:8]
	// Sl_angka = append(Sl_angka, 11, 12, 13)
	// fmt.Println(Sl_angka)
	/*
		Soal 3
		Buatlah sebuah map yang menyimpan informasi tentang buku
		(judul sebagai key, dan jumlah halaman sebagai value).
		Tambahkan 5 buku ke dalam map, lalu hitung dan tampilkan
		jumlah halaman total dari semua buku yang ada.
	*/
	buku := map[string]int{
		"go":    100,
		"js":    120,
		"react": 300,
		"py":    200,
		"rust":  200,
	}

	jumlahHal := 0
	for _, hal := range buku {
		jumlahHal += hal
	}
	fmt.Println(jumlahHal)
}
