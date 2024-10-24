package main

import "fmt"

func sayHelloTo(firtsName string, midleName string, lastName string) {
	fmt.Println("hello", firtsName, midleName, lastName)
}

func urutAngka(angkaSatu int, angkaDua int) {
	fmt.Println(angkaSatu, angkaDua)
}

func main() {
	sayHelloTo("Muhammad", "Torikal", "Akbar")
	urutAngka(3, 4)
}
