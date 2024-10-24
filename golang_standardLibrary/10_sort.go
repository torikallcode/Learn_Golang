package main

import (
	"fmt"
	"sort"
)

/*
// mengurutka struct berdasarkan age

	type User struct {
		Name string
		Age  int
	}

// kontrak
type UserSlice []User

	func (s UserSlice) Len() int {
		return len(s)
	}

	// Mengurutkan berdasarkan umur
	func (s UserSlice) Less(i, j int) bool {
		return s[i].Age < s[j].Age
	}
	// Mengurutkan berdasarkan name
	func (s UserSlice) Less(i, j int) bool {
		return s[i].Name < s[j].Name
	}

	func (s UserSlice) Swap(i, j int) {
		s[i], s[j] = s[j], s[i]
	}
*/
func main() {
	/*
	 */
	// Mengurutkan struct berdasarkan age
	users := []User{
		{"Muhammad", 19},
		{"Torikal", 20},
		{"Akbar", 21},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
	/*
		// mengurutkan int dari yang terkecil
		angka := []int{42, 23, 34, 5, 67}
		sort.Ints(angka)
		fmt.Println(angka)

		// mengurutkan string dari A - z
		buah := []string{"apel", "mangga", "belimbing", "semangka"}
		sort.Strings(buah)
		fmt.Println(buah)

		// Mengurutkan float dari yang terkecil
		desimal := []float64{1.23, 0.23, 1.30, 1.00, 2.33}
		sort.Float64s(desimal)
		fmt.Println(desimal)

		// mengurutkan slice dari yang terbesar
		sort.Slice(angka, func(i, j int) bool {
			return angka[i] > angka[j]
		})
		fmt.Print(angka)
	*/

}
