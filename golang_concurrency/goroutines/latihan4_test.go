/*
Fungsi Anonim dengan Goroutine:
Buat program di mana goroutine menjalankan fungsi anonim
dan coba untuk mengirimkan parameter ke dalamnya.
*/
package goroutines

import (
	"fmt"
	"testing"
)

func TestLatihanEmpat(t *testing.T) {
	hello := func(name string) {
		fmt.Println("hello", name)
	}
	go hello("Akbar")
}
