package main

import (
	"fmt"

	helloWorld "github.com/torikallcode/module-helloWorld/v2"
)

func main() {
	hello := helloWorld.HelloWorld("Akbar")
	fmt.Print(hello)
}
