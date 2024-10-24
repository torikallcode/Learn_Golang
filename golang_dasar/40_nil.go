package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("data kosog")
	} else {
		fmt.Println(data["Name"])
	}
}