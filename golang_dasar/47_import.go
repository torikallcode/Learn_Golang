package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)


func main() {
	name := helper.SayHello("Torikal")
	fmt.Println("Torikal")
	fmt.Println(name)

	// fmt.Println(helper.version)
	fmt.Println(helper.Aplication)
}