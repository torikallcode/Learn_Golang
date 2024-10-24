package main

import "fmt"

// func determinan(n1, n2 [2]float64) float64 {
// 	result := (n1[0] * n2[1]) - (n1[1] * n2[0])
// 	return result
// }

func determinan(e [2][2]float64) float64 {
	result := (e[0][0] * e[1][1]) - (e[0][1] * e[1][0])
	return result
}

// func adj(n1, n2 [2]float64) string {
// 	b11 := n1[0]
// 	b11 = n2[1]
// 	b12 := n1[1] * (-1)
// 	baris1 := [2]float64{b11, b12}

// 	b22 := n2[1]
// 	b22 = n1[0]
// 	b21 := n2[0] * (-1)
// 	baris2 := [2]float64{b21, b22}

// 	return fmt.Sprintf("%v\n%v", baris1, baris2)
// }

func adj(e [2][2]float64) [2][2]float64 {
	var adje [2][2]float64
	adje[0][0] = e[1][1]  // Menukar elemen diagonal utama
	adje[0][1] = -e[0][1] // Menukar tanda
	adje[1][0] = -e[1][0] // Menukar tanda
	adje[1][1] = e[0][0]  // Menukar elemen diagonal utama
	return adje
}

// func invers(d float64, b1, b2 [2]float64) string {
// 	if d == 0 {
// 		fmt.Println("Matriks anda tidak memiliki invers karena nilai determinan anda 0")
// 	}
// 	b11 := 1 / d * b1[0]
// 	b12 := 1 / d * b1[1]
// 	b21 := 1 / d * b2[0]
// 	b22 := 1 / d * b2[1]
// 	baris1 := [2]float64{b11, b12}
// 	baris2 := [2]float64{b21, b22}
// 	return fmt.Sprintf("%.2f\n%.2f", baris1, baris2)
// }

func invers(d float64, e [2][2]float64) string {
	var inve [2][2]float64
	det := 1 / d
	if d == 0 {
		fmt.Println("Invers matriks tidak ada karena determinan = 0")
	}
	for i := 0; i < len(e); i++ {
		for j := 0; j < len(e); j++ {
			inve[i][j] += (e[i][j] * det)
		}
	}
	return fmt.Sprintf("hasil: %2.f", inve)
}

func main() {
	// var baris_satu, baris_dua [2]float64
	// fmt.Print("Masukkan baris pertama: ")
	// for i := 0; i < 2; i++ {
	// 	fmt.Scan(&baris_satu[i])
	// }
	// fmt.Print("Masukkan baris kedua: ")
	// for i := 0; i < 2; i++ {
	// 	fmt.Scan(&baris_dua[i])
	// }

	// fmt.Println("determinan dari matriks: ")
	// fmt.Println(baris_satu)
	// fmt.Println(baris_dua)
	// resultDeterminan := determinan(baris_satu, baris_dua)
	// fmt.Println("=", resultDeterminan)
	// fmt.Println(adj(baris_satu, baris_dua))
	// fmt.Println(invers(resultDeterminan, baris_satu, baris_dua))

	var matriks [2][2]float64

	fmt.Println("Masukkan matriks anda: ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Elemen [%d][%d]: ", i+1, j+1)
			_, err := fmt.Scan(&matriks[i][j])
			if err != nil {
				fmt.Println("Log fatal")
			}
		}
	}
	resultDet := determinan(matriks)
	resultaAdj := adj(matriks)
	fmt.Println(resultDet)
	fmt.Println(resultaAdj)
	fmt.Println(invers(resultDet, resultaAdj))

}

// matriks ordo 2x2
/*
pengguna akan melakukan input untuk setiap baris berupa dua angka
determinan = ad- bc
invers = A -1 = (Adj A)/(det A)
*/
