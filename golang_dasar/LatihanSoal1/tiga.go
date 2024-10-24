// Bangun ruang berbentuk balok dengan panjang p, lebar l, dan tinggi t,
// diketahui memiliki volume sebesar p l t dan luas  permukaan sebesar 2pl + 2pt + 2lt
// Buatlah sebuah program yang menerima input panjang, lebar, tinggi sebuah balok dan mencetak volume dan luasnya pada layar.

package main

import "fmt"

func volumeBalok(panjang , lebar , tinggi float64) float64 {
	return panjang * lebar * tinggi
}

func luasBalok(panjang , lebar , tinggi float64) float64 {
	return (2 * panjang * lebar) + (2 * panjang * tinggi) + (2 * lebar * tinggi)
}

func main() {
	// panjang := 5
	// lebar := 7
	// tinggi := 10

	// volume := panjang * lebar * tinggi
	// luas := (2 * panjang * lebar) + (2* panjang * tinggi) + (2 * lebar * tinggi)

	// fmt.Println("Volume balok= " , volume)
	// fmt.Println("Luas permukaan balok= " , luas)

	fmt.Println(volumeBalok(5, 7 ,9))
	fmt.Println(luasBalok(5,7,9))
}