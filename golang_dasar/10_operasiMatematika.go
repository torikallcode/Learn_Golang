package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := 15
	e := 2
	d := a + b - c * e
	fmt.Println(d)

	f := 10
	f +=5
	fmt.Println(f)

	g := 15
	g++
	fmt.Println(g)
	g--
	fmt.Println(g)
}