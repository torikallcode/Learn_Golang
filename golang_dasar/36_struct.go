package main

import "fmt"

type developer struct {
	Name, Address string
	Age           int
}

func main() {

	var Muhammad developer
	Muhammad.Name = "Muhammad"
	Muhammad.Address = "Indonesia"
	Muhammad.Age = 20

	fmt.Println(Muhammad)

	Torikal := developer{
		Name:    "Torikal",
		Address: "Indonesia",
		Age:     20,
	}

	fmt.Println(Torikal)
	fmt.Println(Torikal.Name)
	fmt.Println(Torikal.Address)
	fmt.Println(Torikal.Age)

	Akbar := developer{"Akbar" , "Indonesia" , 20}
	fmt.Println(Akbar)
}