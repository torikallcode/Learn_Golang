package main

import "fmt"

type Adress struct {
	City, Province, Country string
}

func main() {
	adress1 := Adress{"Selong", "Lotim", "Ntb"}
	adress2 := &adress1

	adress2.City = "Masbagik"
	fmt.Println(adress1)
	fmt.Println(adress2)

	*adress2 = Adress{"Jakarta" , "Jawa Barat" , "Indonesia"}
	fmt.Println(adress1)
	fmt.Println(adress2)
}