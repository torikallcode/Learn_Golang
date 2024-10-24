package main

import "fmt"

func main() {
	person := map[string]string {
		"name" : "Torikal",
		"address" : "Pancor",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := map[string]string{
		"title" : "Bumi",
		"penulis" : "hehe",
		"halaman" : "127",
	}

	fmt.Println(book)
	delete(book , "title")
	fmt.Println(book)

	angka := map[int32]int32{
		1 : 90,
		2: 95,
		3 : 100,
	}

	fmt.Println(angka)
	delete(angka , 1)
	fmt.Println(angka)

}