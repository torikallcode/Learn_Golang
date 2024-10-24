package channels

import (
	"fmt"
	"testing"
)

/*
Buat program yang menggunakan bufferd channel dengan kapasitas 3,
kirimkan 3 data secara berurutan dan terima datanya setelah semuanya
dikirimkan . Lihat bagaimana buffer mempengaruhi perilaku program
*/

func TestBufferedChannel(t *testing.T) {
	ch1 := make(chan int, 3) //membuat buffered channel dengan kapasitas 3
	ch1 <- 1                 //mengirim data 1 ke buffered channel
	ch1 <- 2                 //mengirim data 2 ke buffered chanel
	ch1 <- 3                 // mengirim data 3 ke buffered channel
	for n := range ch1 {
		fmt.Println(n) //mencetak ch1 dengan menggunakan range
	}
}
