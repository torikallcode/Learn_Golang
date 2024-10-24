package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Angka ke" , counter)
		counter++
	}

	for counter := 1; counter <= 10 ; counter ++ {
		fmt.Println("Angka ke-" , counter)
	} 



	names := []string {"Muhammad" , "Torikal" , "Akbar"}

		for i := 0 ; i < len(names); i++ {
			fmt.Println(names[i])
		}

	for index , name := range names {
		fmt.Println("index" , index , "=" , name)
	}


	for counter = 0; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index ke", index , "=" , name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}