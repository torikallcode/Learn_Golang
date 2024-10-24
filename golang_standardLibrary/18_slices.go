package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Muhammad", "Torikal", "Akbar"}
	values := []int{100, 90, 95}

	fmt.Println(slices.Min(names))                 // Mencari nama yang awal hurufnya paling kecil
	fmt.Println(slices.Min(values))                // mencari angka yang paling kecil
	fmt.Println(slices.Contains(names, "Torikal")) // mencari nama di variable names dengan mengembalikan sebuah boolean
	fmt.Println(slices.Contains(values, 1))        // mencari angka di variable values dengan mengembalikan sebuah boolean
	fmt.Println(slices.Index(names, "Akbar"))      // mendapatkan index keberaba dari nama "Akbar"
	fmt.Println(slices.Index(names, "Leal"))
}
