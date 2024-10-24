/*
2. Menggunakan context.WithCancel()
context.WithCancel() membuat Context baru yang bisa dibatalkan secara manual.
Ini berguna jika Anda ingin menghentikan operasi ketika suatu kondisi terpenuhi.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//membuat context dengan pembatalan
	ctx, cancel := context.WithCancel(context.Background())

	//memanggil fungsi yang menerima context
	go doWork(ctx)

	//menunggu beberapa detik sebelum membatalkan
	time.Sleep(3 * time.Second)
	cancel()                    //membatalkan context
	time.Sleep(1 * time.Second) //memberi waktu untuk goeoutine menyelesaikan
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //jika context dibatalkan
			fmt.Println("Pekerjaan dibatalkan")
			return
		default:
			fmt.Println("Sedang bekerja...")
			time.Sleep(1 * time.Second)
		}
	}
}
