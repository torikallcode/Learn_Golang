/*
1. context.Background()
context.Background() adalah titik awal yang paling sering digunakan untuk membuat Context utama.
Biasanya digunakan sebagai root context dalam aplikasi.
*/

package main

import (
	"context"
	"fmt"
)

func main() {
	//membuat context utama
	ctx := context.Background()

	//menggunakan context (misalnya mengirim ke fungsi lain)
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println("Context digunakan dalam fungsi")
}
