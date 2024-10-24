package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&num1)
	fmt.Print("Masukkan operator(+ , - , x , /): ")
	fmt.Scanln(&operator)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("%v + %v = %v", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%v - %v = %v", num1, num2, num1-num2)
	case "x":
		fmt.Printf("%v x %v = %v", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%v / %v = %v", num1, num2, num1/num2)
	default:
		fmt.Println("Masukkan sebuah angka")
	}
}
