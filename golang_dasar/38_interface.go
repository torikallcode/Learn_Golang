package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello" , value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

type Ceo struct {
	Name string
}

func (ceo Ceo) GetName() string {
	return ceo.Name
}

func main() {
	person := Person{Name: "Akbar"}
	SayHello(person)

	ceo := Ceo{"Torikal"}
	SayHello(ceo)
}