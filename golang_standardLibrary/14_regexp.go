package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
		regex := regexp.MustCompile(`e[a-z]o`) //membuat pola
		fmt.Println(regex.MatchString("eyo"))  // cek apakah cocok dengan pola tersebut
		fmt.Println(regex.MatchString("eyi"))
		fmt.Println(regex.MatchString("eYo"))

		//cek apakah cocok dengan pola tersebut: type slice
		fmt.Println(regex.FindAllString("ego , egi , eto , eyo , eyu , elo , epo", 10))

		regex2 := regexp.MustCompile("dunia")         //pola untuk mencari kata dunia
		fmt.Println(regex2.FindString("Hallo dunia")) //Output: dunia

		regex3 := regexp.MustCompile(`[0-9]+`) //pola untuk memecah jika ada angak 0 - 9
		hasil := regex3.Split("abc123def234ghi", -1)
		fmt.Println("hasil split:", hasil) //output:[abc def ghi]
	*/
	// latihansoal mjth
	yourEmail := "Torikal@gmail.com"
	valid, err := regexp.MatchString("@gmail.com", yourEmail)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(valid)
	}
}
