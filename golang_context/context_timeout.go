/*
3. Menggunakan context.WithTimeout()
context.WithTimeout() membuat Context yang secara otomatis dibatalkan setelah batas waktu tertentu.
Sangat bermanfaat untuk operasi yang memiliki batas waktu.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go doWorkTimeout(ctx)

	<-ctx.Done()
	fmt.Println("Context dibatalkan: ", ctx.Err())
}

func doWorkTimeout(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Pekerjaan dihentikan")
			return
		default:
			fmt.Println("Sedang bekerja...")
			time.Sleep(1 * time.Second)
		}
	}
}
