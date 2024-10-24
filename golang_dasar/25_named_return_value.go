package main

import "fmt"

func completeName() (firtsName, midleName, lastName string) {
	firtsName = "Muhammad"
	midleName = "Torikal"
	lastName = "Akbar"

	return firtsName, midleName, lastName
}

func main() {
	a, b, c := completeName()
	fmt.Println(a, b ,c)
}