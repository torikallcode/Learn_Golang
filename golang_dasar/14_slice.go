package main

import "fmt"

func main() {
	nama := [...]string {"Muhammad" , "Torikal" , "Akbar" , "Aisyah" , "Zahra",}

	sliceNama1 := nama[1:4]
	sliceNama2 := nama[2:]
	sliceNama3 := nama[:3]
	sliceNama4 := nama[:]

	fmt.Println(nama)
	fmt.Println(sliceNama1)
	fmt.Println(sliceNama2)
	fmt.Println(sliceNama3)
	fmt.Println(sliceNama4)

	days := [...]string{"senin" , "selasa" , "rabu" , "kamis" , "jumat" , "sabtu" , "minggu"}

	daySlice1 := days[4:]
	daySlice1[1] = "sabtu baru"
	daySlice1[2] = "minggu baru"
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1 , "Libur baru")
	fmt.Println(daySlice2)
	fmt.Println(days	)
}