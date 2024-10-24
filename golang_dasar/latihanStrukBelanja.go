package main

import "fmt"

func main() {
	fmt.Println("---Terimakasih---")
	fmt.Println("Friday , 27 July 2024")
	fmt.Println("_________________________")

	var (
		barang1 = "Coffe"
		barang2 = "Snack"
		barang3 = "Rise"
	)

	fmt.Println("1.",barang1, " = " , "Rp", 5000)
	fmt.Println("2.",barang2, " = " , "RP", 7500)
	fmt.Println("3.",barang3 , " = " , "Rp" , 20000)
	fmt.Println("_________________________")

	fmt.Println("Total = ", 5000 + 7500 + 20000)
}