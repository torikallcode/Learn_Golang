package main

import (
	"fmt"
	"net/http"
)

// import (
// 	"fmt"
// 	"net/http" //package untuk membuat server web dan menangani permnitaan http dari browser
// )

// // Fungsi ini memiliki dua parameter: w (untuk menulis respons ke klien) dan r (untuk membaca permintaan dari klien).
// func helloWorldPage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World") //menuliskan teks "Hello World" ke dalam respons yang dikirim ke browser
// }

// func main() {
// 	http.HandleFunc("/", helloWorldPage) //Baris ini memberitahu server bahwa setiap kali ada permintaan ke URL root ("/"), fungsi helloWorldPage harus dijalankan untuk menampilkan respons.
// 	http.ListenAndServe("", nil)
// 	/*
// 		Baris ini menjalankan server web dan membuatnya siap menerima permintaan dari klien.
// 		Parameter pertama adalah alamat dan port tempat server berjalan. Jika dibiarkan kosong (""), server akan berjalan di port default (biasanya 8080).
// 		Parameter kedua (nil) adalah untuk konfigurasi tambahan, yang dalam kasus ini tidak digunakan.
// 	*/
// }

func helloClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello CLient")
}

func main() {
	http.HandleFunc("/", helloClient)
	http.ListenAndServe("", nil)
}
