package main

import "fmt"

func main() {
	// konversi integer
	// nilai dari nilaiInt16 tidak sesuai karena nilai 32768 melebihi batas int16
	nilaiInt32 := 32768
	nilaiInt64 := int64(nilaiInt32)
	nilaiInt16 := int16(nilaiInt32)
	
	fmt.Println(nilaiInt32)
	fmt.Println(nilaiInt64)
	fmt.Println(nilaiInt16)

	// Konversi string
	// e akan memunculkan byte dari Huruf "T"
	// eString akan konversi byte dari e
	midleName := "Torikal"
	e := midleName[0]
	eString := string(e)

	fmt.Println(midleName)
	fmt.Println(e)
	fmt.Println(eString)
}