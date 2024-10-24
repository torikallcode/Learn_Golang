package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))  //Membulatkan float64 keatas
	fmt.Println(math.Floor(1.60)) //Membulatkan float64 kebawah
	fmt.Println(math.Round(1.49)) //Membualatkan float64 sesuai dengan yang paling dekat
	fmt.Println(math.Max(10, 11)) //mengembalikan nilai float64 paling besar
	fmt.Println(math.Min(10, 11)) // mengembalikan nilai float64 paling kecil

}

// Package math merupakan package yang berisikan constant dan fungsi matematika
