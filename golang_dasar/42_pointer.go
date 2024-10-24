package main

import "fmt"

type Adress struct {
	City, Province, Country string
}

func main() {
	// // Pass by value
	// adress1 := Adress{"Selong", "Lombok timur", "NTB"}
	// adress2 := adress1

	// adress2.City = "Masbagik"
	// fmt.Println(adress1) //Tidak berubah
	// fmt.Println(adress2) //Berubah menjadi masbagik

	// pointer
	adress1 := Adress{"Selong" , "Lombok timur" , "NTB"}
	adress2 := &adress1

	adress2.City = "Masbagik"
	fmt.Println(adress1) //Ikut berubah
	fmt.Println(adress2) //berubah menjadi masbagik

}