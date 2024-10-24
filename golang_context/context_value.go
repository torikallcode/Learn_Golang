/*
4. Menggunakan context.WithValue()
context.WithValue() memungkinkan Anda menambahkan nilai tambahan dalam Context.
Hal ini digunakan untuk menyimpan data seperti ID pengguna atau informasi otentikasi.
*/

package main

import (
	"context"
	"fmt"
)

func main() {
	//membuat context dengan nilai
	ctx := context.WithValue(context.Background(), "userID", 42)

	//memanggil fungsi dan meneruskan context dengan nilai
	printUserID(ctx)
}

func printUserID(ctx context.Context) {
	// mengambil nilai dari context
	userID := ctx.Value("userID")
	fmt.Printf("User ID: %v\n", userID)
}
