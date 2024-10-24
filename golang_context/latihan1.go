package main

import (
	"context"
	"fmt"
	"time"
)

/*
Latihan 1: Membuat Pembatalan Manual
Buatlah program yang menggunakan context.WithCancel(). Jalankan goroutine yang menampilkan pesan "Sedang bekerja..."
setiap detik, dan hentikan setelah 5 detik dengan memanggil fungsi pembatalan.
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go batalkan(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func batalkan(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Pekerjaan dibatalkan")
			return
		default:
			fmt.Println("Sedang bekerja...")
			time.Sleep(1 * time.Second)
		}
	}
}
