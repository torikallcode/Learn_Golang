package main

import "fmt"

type Adress struct {
	City, Province, Country string
}

func main() {
	alamat1 := new(Adress)
	alamat2 := alamat1

	alamat2.City = "Selong"
	alamat2.Province = "Lotim"
	alamat2.Country = "NTB"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}