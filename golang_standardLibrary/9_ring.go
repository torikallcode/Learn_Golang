package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// membuat ring baru
	var data *ring.Ring = ring.New(5)

	// membuat program untuk memanggil elemen ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.Itoa(i)
		data = data.Next()
	}

	// data.Value = "value 1"

	// data = data.Next()
	// data.Value = "value 2"

	// data = data.Next()
	// data.Value = "value 3"

	// data = data.Next()
	// data.Value = "value 4"

	// data = data.Next()
	// data.Value = "value 5"

	// memanggil elemen ring
	data.Do(func(a any) {
		fmt.Println(a)
	})

}
