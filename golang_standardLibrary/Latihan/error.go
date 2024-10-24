package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("Validation error")
	notFoundError   = errors.New("NotFound Error")
)

func getName(n string) error {
	if n == "" {
		return validationError
	}

	if n != "Akbar" {
		return notFoundError
	}

	return nil
}

func main() {

	nama := ""

	err := getName(nama)
	fmt.Println(err)
	// if err != nil {
	// 	if errors.Is(err, validationError) {
	// 		fmt.Println("Validation Error")
	// 	} else if errors.Is(err, notFoundError) {
	// 		fmt.Println("Not Found Error")
	// 	} else {
	// 		fmt.Println("Unknown error")
	// 	}
	// } else {
	// 	fmt.Printf("Nama anda adalah: %s", nama)
	// }
}
