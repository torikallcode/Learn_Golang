package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	counter int        //variable yang akan ditambahkan
	mu      sync.Mutex //digunakan untuk mengunci counter sehingga hanya ada satu goroutine yang dapat mengubah counter dalam satu waktu
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done() //menandakan goroutine selesai

	mu.Lock() //mengunci counter agar hanya 1 goroutine yang bisa akses
	counter++
	fmt.Printf("Counter incremented: %d\n", counter)
	time.Sleep(1 * time.Second)
	mu.Unlock() //membuka kunci agar goroutine lain bisa akses
}

func TestMutext(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg) //menalankan 5 goroutine yang mengakses counter
	}

	wg.Wait() //menunggu sampai semua gorutine selesai
	fmt.Println("Final Counter:", counter)
}
