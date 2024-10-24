package selectstatement

import (
	"fmt"
	"testing"
	"time"
)

// func TestSelectStatement(t *testing.T) {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	// Goroutine pertama
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch1 <- "Data dari ch1"
// 	}()

// 	// Goroutine kedua
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch2 <- "Data dari ch2"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println("Menerima:", msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println("Menerima:", msg2)
// 		}
// 	}
// }

func RataRata(angka []int, ch1 chan int) {
	result := 0
	for _, n := range angka {
		result += n
		time.Sleep(1 * time.Second)
	}
	ch1 <- result / len(angka) //ch1 menerima data berupa result
}

func Max(angka []int, ch1 chan int) {
	max := angka[0]
	for _, n := range angka {
		if max < n {
			max = n
		}
		time.Sleep(1 * time.Second)
	}
	ch1 <- max //ch1 menerima data berupa max
}

func TestSelectStatement(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	n := []int{1, 2, 3, 4, 5}
	fmt.Println("Angka:", n)
	go RataRata(n, ch1) //memanggil fungsi RataRata dengan memasukkan slice n dan channels ch1
	go Max(n, ch2)      //memanggil fungsi Max dengan memasukkan slice n dan channels ch2

	for i := 0; i < 2; i++ { //melakukan perulangan untuk kedua fungsi dan channels
		select {
		case rtr := <-ch1: // menerima data dari channels ch1
			fmt.Println("Rata rata:", rtr)
		case max := <-ch2: //menerima data dari channels ch2
			fmt.Println("Max:", max)
		default:
			fmt.Println("Tidak ada data yang diterima")
		}
	}
	defer close(ch1)
	defer close(ch2)

}
