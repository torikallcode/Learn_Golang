/*
Eksperimen dengan Goroutines:
Coba buat program yang menjalankan 100 goroutines yang mencetak
angka secara bersamaan. Gunakan sync.WaitGroup untuk menunggu
hingga semua goroutines selesai.
*/

package goroutines

import (
	"sync"
	"testing"
)

// func angkan(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(i)
// 	}
// }

func TestLatihanTiga(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		for i := 0; i < 100; i++ {
			go angka()
		}
	}
}
