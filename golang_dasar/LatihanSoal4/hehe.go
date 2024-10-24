package main

import "fmt"

func main() {
	nilai := []int{
		56, 76, 85, 90, 95, 67, 88, 32, 25,
	}
	// rata rata
	rata_rata := 0
	for _, value := range nilai {
		rata_rata += value
	}
	fmt.Println(rata_rata)

	// nilai minimum
	min_nilai := nilai[0]

	for _, value := range nilai {
		if value < min_nilai {
			min_nilai = value
		}
	}
	fmt.Println(min_nilai)
}
