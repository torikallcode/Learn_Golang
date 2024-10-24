package main

import "fmt"

func main() {
	var nilai1 = 95
	var nilai2 = 100

	var lulusNilai1 bool = nilai1 > 85
	var lulusNilai2 bool = nilai2 > 85

	var lulus bool = lulusNilai1 && lulusNilai2

	fmt.Println(lulus)
}