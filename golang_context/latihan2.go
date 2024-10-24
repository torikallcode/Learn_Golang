/*
Latihan 2: Menggunakan context.WithTimeout()
Implementasikan program yang melakukan pekerjaan dalam goroutine selama 10 detik,
tetapi batasi waktu eksekusi menggunakan context.WithTimeout() selama 3 detik.
Tampilkan pesan ketika pekerjaan selesai atau ketika dibatalkan karena batas waktu.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//membuat context dengan batas waktu 3 detik
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() //membatalkan ketika selesai

	go angka(ctx)

	//jika context dibatalkan
	<-ctx.Done()
	fmt.Println("Context dibatalkan", ctx.Err())
}

func angka(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //jika context dibatalkan
			fmt.Println("Perhitungan selesai")
			return
		default: //program yang dijalankan ketika context belum dibatalkan
			for i := 1; i < 10; i++ {
				fmt.Println(i)
				time.Sleep(1 * time.Second)
			}
		}
	}
}
