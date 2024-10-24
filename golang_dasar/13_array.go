package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Torikal"
	names[2] = "Akbar"

	fmt.Println(names)

	var value = [...]int {
		90 , 95 , 100,
	}

	fmt.Println(value)
	value[2] = 120
	fmt.Println(value)
	fmt.Print(len(value))

}