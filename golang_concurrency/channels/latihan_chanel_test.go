package channels

import (
	"fmt"
	"testing"
)

func Angka(ch1 chan int) {
	for i := 0; i <= 10; i++ {
		ch1 <- i
	}
	defer close(ch1)
}

func TestAngka(t *testing.T) {
	ch1 := make(chan int)
	go Angka(ch1)
	for n := range ch1 {
		fmt.Println(n)
	}
}
