package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) SayHello(name string){
	fmt.Println("Hello" , name ,"My name is" , customer.Name )
}

func main() {
	Torikal := Customer{
		Name:    "Torikal",
		Address: "Indonesia",
		Age:     20,
	}

	fmt.Println(Torikal)
	Torikal.SayHello("Akbar")


}