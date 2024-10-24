package channels

import (
	"fmt"
	"testing"
	"time"
)

/*
Buat program yang menggunakan bufferd channel dengan kapasitas 5,
kirimkan 10 angka secara berurutan ke channel dan terima angka tersebut secara
perlahan dengan delay 1 detik untuk setiap penerimaan,
lihat bagaimana buffer menyimpan data sementara penerima tidak segera
mengambil data
*/

func TestBufferedChannel2(t *testing.T) {
	ch1 := make(chan int, 5) //membuat variable untuk menampung channel dengan kapasitas 5
	for i := 1; i <= 5; i++ {
		ch1 <- i //memasukkan 5 nilai untuk buffered channel dengan kapasitas 5
	}
	close(ch1) //close channel

	for n := range ch1 {
		fmt.Println("Menerima:", n) //mencetak variable ch1
		time.Sleep(1 * time.Second) //delay 1 detik setiap mencetak 1 channel
	}

	fmt.Println("Selesai")

}
