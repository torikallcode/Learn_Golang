package main

import "fmt"

func getHalo(name string) string {
	return "Halo " + name
}

func main() {
	contoh := getHalo

	fmt.Println(contoh("Akbar"))
}