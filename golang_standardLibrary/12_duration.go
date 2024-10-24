package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second //mengambil 100 detik
	duration2 := 100 * time.Millisecond             //mengambil 100 milidetik
	duration3 := duration1 - duration2              // 100 detik - 100 milidetik , output: 1m39.9s
	duration4 := 1 * time.Hour                      //mengambil 1 jam
	duration5 := 10 * time.Minute                   // mengambil 10 menit
	fmt.Println(duration1)
	fmt.Printf("duration: %d\n", duration1)
	fmt.Println(duration3)
	fmt.Println(duration4)
	fmt.Println(duration5)
}
