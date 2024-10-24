package main

import (
	"fmt"
	"time"
)

func main() {
	// var sekarang time.Time = time.Now()
	// fmt.Println(sekarang)

	// var utc time.Time = time.Date(2003, time.December, 13, 0, 0, 0, 0, time.UTC)
	// fmt.Println(utc)
	// fmt.Println(utc.Local())

	// formatter := "2006-01-02 15:04:05"
	// value := "2024-09-18 08:52:01"
	// valueTime, err := time.Parse(formatter, value)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// } else {
	// 	fmt.Println(valueTime)
	// }

	// fmt.Println(valueTime.Hour())

	// soal mjth
	// menambahkan 10 hari kedepan
	waktu_sekarang := time.Now()
	format_waktu := "2006-01-02 15:04:05"
	sepuluhHari := waktu_sekarang.AddDate(0, 0, 10)
	fmt.Println(sepuluhHari.Format(format_waktu))

}
