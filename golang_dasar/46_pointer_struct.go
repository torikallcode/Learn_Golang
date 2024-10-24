package main

import "fmt"

type Data struct {
	Name string
}

func (data *Data) meried(){
	data.Name = "Mr. " + data.Name
}

func main() {
	Torikal := Data{"Torikal"}
	Torikal.meried()

	fmt.Println(Torikal.Name)
}