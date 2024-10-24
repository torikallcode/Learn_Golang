package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var point [7]int
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	var tebakan int
	for i := 0; i < len(point); i++ {
		fmt.Print("Tebak angka: ")
		fmt.Scanln(&tebakan)
		if tebakan == random {
			fmt.Printf("Selamat anda berhasil menebak angka: %v", random)
			break
		} else if tebakan > random {
			fmt.Println("Tebakan anda salah, angka yang anda masukkan lebih besar")
		} else if tebakan < random {
			fmt.Println("Tebakan anda salah, Angka yang anda masukkan lebih kecil")
		}
	}
}
