package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Muhammad, Torikal , Akbar\n" +
		"Torikal , Akbar , leal\n" +
		"Akbar , leal , Torikal"

	reader := csv.NewReader(strings.NewReader(csvString)) //membaca dari variable csvString
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
