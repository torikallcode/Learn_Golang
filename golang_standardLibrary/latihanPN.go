package main

import "fmt"

func posnes(numbers []int) []int {

	if len(numbers) == 0 {
		return []int{}
	}

	positif := 0
	negatif := 0
	for _, num := range numbers {
		if num > 0 {
			positif++
		} else if num < 0 {
			negatif += num
		}
	}
	return []int{positif, negatif}
}
func main() {
	angka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Print(posnes(angka))
}

// Untuk masukan [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], kamu harus kembali [10, -65].
// package kata

// func CountPositivesSumNegatives(numbers []int) []int {
//   res := []int{0,0}
//   for _, v := range numbers {
//     if v > 0 {
//       res[0] += 1
//     }else {
//       res[1] += v
//     }
//   }
//   return res // your code here
// }
