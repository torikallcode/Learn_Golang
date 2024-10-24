package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func Number() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func Ganjil() {
	for i := 1; i <= 9; i += 2 {
		fmt.Println(i)
	}
}

func Genap() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func TestAnonim(t *testing.T) {
	hello := func(name string) {
		fmt.Println("Hello", name)
	}
	go hello("Akbar")
	go Number()
	go Ganjil()
	go Genap()

	fmt.Println("UPS")

	time.Sleep(1 * time.Second)
}
