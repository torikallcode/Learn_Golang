package channels

import (
	"fmt"
	"testing"
	"time"
)

func TestChannels(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello world"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestChannelsInt(t *testing.T) {
	number := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		number <- 13
		fmt.Println("Hello")
	}()

	value := <-number
	fmt.Println("angka dari variable number:", value)
	fmt.Println("Done")
	time.Sleep(5 * time.Second)
}
