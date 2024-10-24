package main

import (
	"fmt"
)

type Filter func(string) string

func sayHellowithFilter(name string, filter Filter ) {
	fmt.Println("Hello " , filter(name))
}

func filterBadWord(word string) string {
	if word == "Anjing"{
		return "..."
	} else {
		return word
	}
}


func main() {
	sayHellowithFilter("akbar", filterBadWord)
	// fmt.Println(sayHellowithFilter)
	filter := filterBadWord
	sayHellowithFilter("Anjing", filter)

	// fmt.Println(modulus(10 , 2))

}