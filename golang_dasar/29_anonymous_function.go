package main

import "fmt"

type BlackList func(string) bool

func register(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("kamu di blok", name)
	} else {
		fmt.Println("Halo" , name)
	}
}

func main() {
	blackList := func(name string) bool {
			return name == "kata kasar"
	}
	register("Akbar" , blackList)


	register("Akbar", func(name string) bool {
		return name == "kata kasar"
	})
}