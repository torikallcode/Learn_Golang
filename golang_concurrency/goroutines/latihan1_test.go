/*
Menjalankan Goroutine Tanpa Menunggu:
Buat program di mana beberapa goroutines dijalankan,
tetapi program utama tidak menunggu mereka selesai.
Perhatikan apa yang terjadi.
*/

package goroutines

import (
	"fmt"
	"testing"
)

func angka() {
	for i := 0; i < 10; i++ {
		fmt.Println("Angka:", i)
	}
}

func TestLatihanSatu(t *testing.T) {
	go angka()
	fmt.Println("ini program utama")
}
