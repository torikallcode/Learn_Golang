/*
Latihan 4: Membatalkan Pekerjaan dengan context.WithCancel()
Buatlah sebuah program yang menjalankan tiga goroutine secara bersamaan, masing-masing melakukan pekerjaan yang berbeda.
Gunakan context.WithCancel() untuk membatalkan semua pekerjaan ketika salah satu pekerjaan selesai lebih dulu.
Tampilkan pesan yang menunjukkan pekerjaan mana yang selesai lebih dulu dan bahwa pekerjaan lainnya dibatalkan.
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(3)

	go doWorkLatihan4(ctx, "Pekerjaan 1", &wg)
	go doWorkLatihan4(ctx, "pekerjaan 2", &wg)
	go doWorkLatihan4(ctx, "pekerjaan 3", &wg)

	go func() {
		defer cancel()
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
		fmt.Println("Salah satu pekerjaan selesai lebih dulu. Membatalkan semua pekerjaan...")
	}()

	wg.Wait()
	fmt.Println("Semua pekerjaan telah selesai dibatalkan")
}

func doWorkLatihan4(ctx context.Context, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s dibatalkan\n", name)
			return
		default:
			fmt.Printf("%s sedang bekerja...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}
