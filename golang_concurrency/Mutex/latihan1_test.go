package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Latihan 1: Sinkronisasi dengan Mutex
Buatlah program yang menjalankan 10 goroutine secara bersamaan.
Setiap goroutine akan menambahkan angka ke variabel bersama,
dan gunakan mutex untuk menghindari data race.

Petunjuk:

Buat variabel bersama untuk menyimpan jumlah (sum).
Setiap goroutine akan menambahkan angka ke variabel tersebut.
Gunakan mutex untuk melindungi akses ke variabel bersama.
*/

var (
	hasil int
	mtx   sync.Mutex
)

func tambah(wg *sync.WaitGroup) {
	defer wg.Done()
	mtx.Lock()
	hasil++
	fmt.Println("Tambah:", hasil)
	time.Sleep(1 * time.Second)
	mtx.Unlock()
}

func TestLatihanMutext(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go tambah(&wg)
	}
	wg.Wait()
	fmt.Println("Hasil akhir tambah", hasil)
}
