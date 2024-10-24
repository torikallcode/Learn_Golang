/*
Latihan 3: Mengirim Nilai dengan context.WithValue()
Buat program yang meneruskan beberapa informasi pengguna (seperti userID, role) menggunakan context.WithValue().
Ambil nilai tersebut dalam fungsi yang berbeda dan tampilkan di layar.
*/

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "userID", 13)
	ctx = context.WithValue(ctx, "role", "Software Engineer")

	printInfoUser(ctx)
}

func printInfoUser(ctx context.Context) {

	useriD, ok := ctx.Value("userID").(int)
	if !ok {
		fmt.Println("User ID tidak ditemukan dan tipe tidak sesuai")
	} else {
		fmt.Printf("User id: %d\n", useriD)
	}

	role, ok := ctx.Value("role").(string)
	if !ok {
		fmt.Println("Role tidak ditemukan dan tipe tidak sesuai")
	} else {
		fmt.Printf("Role: %s\n", role)
	}
}
