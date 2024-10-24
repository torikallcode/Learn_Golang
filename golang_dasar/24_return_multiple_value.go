package main

import "fmt"

func namaMu() (string, string) {
	return "Torikal", "Akbar"
}

func main() {
	// firtsName, lastName := namaMu()
	// fmt.Println(firtsName, lastName)

	firtsName, _ := namaMu()
	fmt.Println(firtsName)
}