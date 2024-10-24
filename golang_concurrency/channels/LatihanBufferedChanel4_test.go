package channels

import (
	"fmt"
	"testing"
	"time"
)

/*
Buat dua goroutines , satu goroutine megirim data ke buffered
channel dan satu lagi menerima data, gunakan buffered channel untuk
mengelola pengiriman data penerimaan secara efisien dan tanpa blockig
yang berlebihan
*/
// goroutine untuk mengirim data ke buffered channel
func send(ch1 chan int) {
	for i := 1; i <= 3; i++ {
		ch1 <- i //memasukkan tiga data kedalam buffered chanel
	}
	close(ch1) //close buffered channel
}

func TestLatihanBuffered4(t *testing.T) {
	ch1 := make(chan int, 3) //membuat variable untuk buffered channel

	go send(ch1) //memanggil goroutine send dengan parameter buffered channel
	for n := range ch1 {
		fmt.Println("Menerima data:", n) //mencetak data yang dikirim ke buffered channel
		time.Sleep(1 * time.Second)      //delay 1 detik setiap mencetak data
	}
	fmt.Println("Selesai menerima data")
}
