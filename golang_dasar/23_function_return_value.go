package main

import "fmt"

func getHello(name string) string {
	selamat := "hallo " + name
	return selamat
}

func tambah(a, b int) int {
	return a + b
}

func main() {
	result := getHello("Akbar")
	fmt.Println(result)
	fmt.Println(getHello("Torikal"))
	fmt.Println(tambah(2, 3))
}
