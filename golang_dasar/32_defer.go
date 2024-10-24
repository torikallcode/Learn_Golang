package main

import "fmt"

func endApp() {
	fmt.Println("Program selesai")
}

func hallo(name string){
	defer endApp()
	fmt.Println("hallo", name)
}

func main(){
	hallo("akbar")
}