// Air dalam bentuk es pada suhu <= 0o Celcius
// Air dalam bentuk cair pada suhu 0o s.d 100o Celcius
// Air dalam bentuk uap pada suhu > 100o Celcius

// Konversi suhu dari skala Fanrenheit (oF)
// ke skala Celcius (oC) dapat dihitung dengan
// rumus C = (F-32) x 5 / 9
// C = nilai suhu dalam Celcius
// F = nilai suhu dalam Fanrenheit

// Buat sebuah program yang menerima suhu dalam skala Fanrenheit kemudian mengkonversinya ke dalam skala Celcius. Cetak bentuk/wujud air pada suhu tersebut.

package main

import "fmt"

func celcius(fahrenheit float32) (float32, string) {
	celciusConv := (fahrenheit - 32) * 5 / 9
	bentuk := ""
	if celciusConv <= 0 {
		bentuk = "Es"
	} else if celciusConv <= 100 {
		bentuk = "cair"
	} else {
		bentuk = "uap"
	}

	return celciusConv , bentuk
}

func main(){
	// fahrenheit := 10.0
	// celcius := (fahrenheit - 32) * 5 / 9

	// bentuk := ""
	// if celcius <= 0 {
	// 	bentuk = "Es"
	// } else if celcius <= 100 {
	// 	bentuk = "cair"
	// } else {
	// 	bentuk = "uap"
	// }


	// fmt.Println(celcius)
	// fmt.Println(bentuk)

	fmt.Println(celcius(100.0))
}