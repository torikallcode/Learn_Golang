package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //menandakan goruotine selesai
	fmt.Printf("Worker %d mulai bekerja\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d selesai bekerja\n", id)
}

func TestWorker(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Semua pekerjaan selesai")
}
