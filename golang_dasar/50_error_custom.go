package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{Message: "Validation error"}
	}

	if id != "Torikal" {
		return &NotFoundError{Message: "data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("Torikal", nil)
	if err != nil {
		// terjadi error
		// if ValidationError, ok := err.(*ValidationError); ok {
		// 	fmt.Println("Validation Error:", ValidationError.Error())
		// } else if NotFoundError, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("Not Found Error:", NotFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *ValidationError:
					fmt.Println("Validation Error:", finalError.Error())
		case *NotFoundError:
					fmt.Println("Not Found Error:", finalError.Error())
		default:
					fmt.Println("Unknown error:", err.Error())
		}
	} else {
		// Sukses
		fmt.Println("Sukses")
	}

}