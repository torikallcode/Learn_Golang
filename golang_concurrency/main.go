package main

import "fmt"

// var counterR int

// func incrementR(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		counterR++
// 	}
// }
func Grow(arr []int) int {
	result := 1
	for _, n := range arr {
		result *= n
	}
	return result
}

func main() {
	// var wg sync.WaitGroup

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go incrementR(&wg)
	// }
	// wg.Wait()
	// fmt.Println("Final counter:", counterR)
	angka := []int{1, 2, 3}
	fmt.Println(Grow(angka))
}
