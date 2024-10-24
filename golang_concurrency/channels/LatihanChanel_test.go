package channels

import (
	"fmt"
	"testing"
)

/*
Buat program yang memiliki goroutine yang mengirimkan 10 angka ke chanel,
dan goroutine utama mencetak angka angka tersebut satu persatu
*/

func nomor(ch1 chan int) { //membuat parameter untuk menerima channels
	for i := 0; i < 10; i++ {
		ch1 <- i //mengirim agka 0 - 10 kedalam channel menggunakan perulangan
	}
	defer close(ch1) //close channel
}

func TestLatihanSatuChanel(t *testing.T) {
	ch1 := make(chan int) //Membuat variable untuk menampung channels

	go nomor(ch1) //memanggil goroutine nomor dengan parameter ch1
	for n := range ch1 {
		fmt.Println(n) //mencetak ch1 dengan for range
	}
}
