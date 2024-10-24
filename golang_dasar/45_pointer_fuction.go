package main

import "fmt"

type Adress struct {
	City, Province, Country string
}

func changeAdress(adress *Adress){
	adress.City = "Selong"
	adress.Province = "NTB"
	adress.Country = "Indonesia"
}

func main() {

	Adress := Adress{}
	changeAdress(&Adress)

	fmt.Println(Adress)

}