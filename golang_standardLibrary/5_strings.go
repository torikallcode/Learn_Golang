package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("M.Torikal Akbar", "Akbar"))               //Mencari kata "Akbar"
	fmt.Println(strings.Split("M. Torikal Akbar", ""))                      //untuk memisah string
	fmt.Println(strings.ToLower("M. Torikal Akbar"))                        //membuat string menjadi huruf kecil semua
	fmt.Println(strings.ToUpper("M. Torikal Akbar"))                        // membuat string menjadi huruf besar semua
	fmt.Println(strings.Trim("  M. Torikal Akbar   ", " "))                 // menghapus spasi pada awal dan akhir string
	fmt.Println(strings.ReplaceAll("Torikal Akbar", "Torikal", "Muhammad")) //mengganti kata "Torikal" dengan "Muhammad"
}

/*
// Package string adalah package yang berisikan function function untuk memanipulasi
tipe data string
	fmt.Println(strings.Contains("M.Torikal Akbar", "Akbar"))               //Mencari kata "Akbar"
	fmt.Println(strings.Split("M. Torikal Akbar", ""))                      //untuk memisah string
	fmt.Println(strings.ToLower("M. Torikal Akbar"))                        //membuat string menjadi huruf kecil semua
	fmt.Println(strings.ToUpper("M. Torikal Akbar"))                        // membuat string menjadi huruf besar semua
	fmt.Println(strings.Trim("  M. Torikal Akbar   ", " "))                 // menghapus spasi pada awal dan akhir string
	fmt.Println(strings.ReplaceAll("Torikal Akbar", "Torikal", "Muhammad")) //mengganti kata "Torikal" dengan "Muhammad"
*/
