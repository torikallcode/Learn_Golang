package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))             //mendapatkan jalur direktori dari sebuah file
	fmt.Println(path.Base("hello/world.go"))            //mendapatkan nama file atau folder terakhir dari sebuah jalur
	fmt.Println(path.Ext("hello/world.go"))             //mendapatkan ectensi file dari sebuah jalur
	fmt.Println(path.Join("hello", "world", "main.go")) //menggabungkan beberapa bagian jalur file mnjadi satu jalur
}
