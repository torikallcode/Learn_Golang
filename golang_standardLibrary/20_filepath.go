package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))   //false: apakah dimulai dari depan misal drive c atau d
	fmt.Println(filepath.IsLocal("hello/world.go")) //true: karena tidak dituliskan dari depan
	fmt.Println(filepath.Join("hello", "wordl", "main.go"))
}
