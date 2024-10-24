/*
Mengelola Waktu Tunggu:
Buat program yang menjalankan beberapa goroutines dan gunakan time.Sleep
untuk memastikan goroutine menyelesaikan tugasnya sebelum program selesai.
*/

package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func urut() {
	for i := 0; i < 10; i++ {
		fmt.Println("angka:", i)
	}
}

func TestLatihanDua(t *testing.T) {
	go urut()
	time.Sleep(1 * time.Second)
	fmt.Println("Progam selesai")
}
