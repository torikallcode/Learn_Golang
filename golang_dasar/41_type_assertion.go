package main

import (
	"fmt"
)

func random() any {
	return "Ok"
}

func main() {

	/*
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)
	// Akan terjadi panic karena tipe data tidak sesuai
	resultInt := result.(int) 
	fmt.Println(resultInt)
	*/

	result := random()

	switch value := result.(type){
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Uknown", value)
	}
	

}