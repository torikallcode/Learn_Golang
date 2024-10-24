package main

import "fmt"

func main() {

	counter := 0

	increment := func() {
		counter += 2
	}
	for i := 0; i < 10; i++ {
		increment()
	}

	fmt.Println(counter)

}
