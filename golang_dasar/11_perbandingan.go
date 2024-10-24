package main

import "fmt"

func main() {
	var name1 = "Torikal"
	var name2 = "Torikal"
	a := 5
	b := 10

	var result = name1 == name2
	var result2 = name1 != name2
	var result3 = a >= b

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
}