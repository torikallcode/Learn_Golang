package main

import "fmt"

func endProgram() {
	fmt.Println("Program selesai")
}

func sayHai(error bool){
	defer endProgram()
	if error {
		panic("program error")
	}
}

func main(){
	sayHai(true)
}