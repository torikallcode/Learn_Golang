package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout) //menampilkan di terminal

	_ = writer.Write([]string{"Muhammad", "Torikal", "Akbar"})
	_ = writer.Write([]string{"Torikal", "Akbar", "Leal"})
	_ = writer.Write([]string{"Akbar", "Leal", "Torikal"})
	writer.Flush()

}
