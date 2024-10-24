package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("Program selesai")
	pesan := recover()
	fmt.Println("Terjadi error yaitu" , pesan)
}

func sayHai(error bool){
	defer endApp()
	if error {
		panic("terjadi error")
	}
}

func main() {
	sayHai(true)
	fmt.Println("Haii!! Torikal")
}