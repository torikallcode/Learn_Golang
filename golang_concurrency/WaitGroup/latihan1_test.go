package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Number(id, n int, wg *sync.WaitGroup) {
	defer wg.Done() // Pastikan WaitGroup berkurang setelah goroutine selesai

	fmt.Printf("Goroutine %d mulai, mencetak angka %d\n", id, n)
	time.Sleep(time.Duration(n) * time.Second) // Simulasi kerja berbeda dengan sleep
	fmt.Printf("Goroutine %d selesai, angka %d\n", id, n)

}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)            // Tambahkan 1 ke WaitGroup untuk setiap goroutine
		go Number(i, i, &wg) // Menjalankan goroutine dengan tugas berbeda
	}

	wg.Wait() // Menunggu semua goroutine selesai
	fmt.Println("Program Ganjil Selesai")
}
