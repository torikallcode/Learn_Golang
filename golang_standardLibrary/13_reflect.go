package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name, Address string
	Age           int
	status        bool
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println(valueType.Name(), ":")
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "=", valueField.Type)
	}
}

func main() {
	readField(Person{"Akbar", "Pancor", 20, true})
	var x = 45
	var y = "hello"
	var z = true

	fmt.Println("Tipe dari x:", reflect.TypeOf(x))
	fmt.Println("Tipe dari y:", reflect.TypeOf(y))
	fmt.Println("Tipe dari z:", reflect.TypeOf(z))
}
