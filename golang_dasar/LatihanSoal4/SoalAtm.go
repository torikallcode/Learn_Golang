package main

import "fmt"

func main() {
	saldo := 600000
	pin := "090522"

	fmt.Println("Selamat Datang di ATM Bank Torikal")
	fmt.Println("Kode: 220602047")
	fmt.Println("Masukkan PIN anda")

	var masukkanPin string
	fmt.Scanln(&masukkanPin)
	if masukkanPin != pin {
		fmt.Println("Maaf, PIN anda salah, kartu anda disita untuk keamanan")
		return
	} else {
		fmt.Println("Saldo anda saat ini sebesar:", saldo)
		fmt.Println("Silahkan pilih transaksi yang anda inginkan")
		fmt.Println("(1)Tarik tunai", "(2)Transfer", "(3)Bayar tagihan", "(4)Batal")
	}

	var pilihanTransaksi int
	fmt.Scanln(&pilihanTransaksi)
	if pilihanTransaksi == 4 {
		fmt.Println("Anda membatalkan transaksi, silahkan ambil kartu anda")
		return
	} else if pilihanTransaksi == 1 {
		fmt.Println("Pilih nominal yang ingin anda tarik")
		fmt.Println("(1)100.000", "(2)500.000", "(3)1.000.000", "(4)Batal")
		var tarikTunai int
		fmt.Scanln(&tarikTunai)
		tunaiSatu := 100000
		tunaiDua := 500000
		tunaiTiga := 1000000

		if tarikTunai == 4 {
			fmt.Println("Anda membatalkan transaksi, silahkan ambil kartu anda")
			return
		} else if tarikTunai == 1 {
			if saldo > tunaiSatu {
				fmt.Println("Silahkan ambil uang dan kartu anda, Saldo anda tersisa", saldo-tunaiSatu)
			} else {
				fmt.Println("Saldo anda tidak cukup, silahkan ambil kartu anda")
				return
			}
		} else if tarikTunai == 2 {
			if saldo > tunaiDua {
				fmt.Println("Silahkan ambil uang dan kartu anda, Saldo anda tersisa", saldo-tunaiDua)
			} else {
				fmt.Println("Saldo anda tidak cukup, silahkan ambil kartu anda")
				return
			}
		} else if tarikTunai == 3 {
			if saldo > tunaiTiga {
				fmt.Println("Silahkan ambil uang dan kartu anda, Saldo anda tersisa", saldo-tunaiTiga)
			} else {
				fmt.Println("Saldo anda tidak cukup, silahkan ambil kartu anda")
				return
			}
		}
	} else if pilihanTransaksi == 2 {
		fmt.Println("Masukkan nomer rekening tujuan: ")
		var norek int
		fmt.Scanln(&norek)
		fmt.Println("Masukkan nominal: ")
		var nominal int
		fmt.Scanln(&nominal)
		fmt.Println("Masukkan pesan: ")
		var pesan string
		fmt.Scanln(&pesan)

		if saldo >= nominal {
			fmt.Println("Anda akan mentransfer dana sejumlah:", nominal, "ke nomer rekening", norek, "dengan pesan", pesan)

			fmt.Println("(1.OK, 2. Batal)")
			var konfirmasi int
			fmt.Scanln(&konfirmasi)

			if konfirmasi == 1 {
				fmt.Println("Transfer berhasil, saldo anda tersisa:", saldo-nominal)
				return
			} else if konfirmasi == 2 {
				fmt.Println("Transfer dibatalkan")
				return
			} else {
				fmt.Println("Input yang anda berikan tidak valid")
				return
			}
		} else {
			fmt.Println("Saldo tidak cukup , silahkan ambil kartu anda")
			return
		}
	} else if pilihanTransaksi == 3 {
		fmt.Println("Pilih tagihan (1)Kartu kredit (2)Telepon: ")
		var tagihan int
		fmt.Scanln(&tagihan)

		if tagihan == 1 {
			fmt.Println("Terjadi kesalahan , fitur tidak dapat digunakan")
			return
		} else if tagihan == 2 {
			fmt.Println("Masukkan nomer telepon anda")
			var noHp string
			fmt.Scanln(&noHp)
			totalTagihan := 250000
			fmt.Println("Anda akan membayar tagihan telepon untuk nomer", noHp, "sejumlah", totalTagihan)
			fmt.Println("(1.OK, 2. Batal)")
			var konfirmasi int
			fmt.Scanln(&konfirmasi)
			if konfirmasi == 1 {
				fmt.Println("Pembayaran berhasil , saldo anda tersisa", saldo-totalTagihan)
				return
			} else if konfirmasi == 2 {
				fmt.Println("Pembayaran dibatalkan")
				return
			} else {
				fmt.Println("Input yang anda berikan tidak valid")
				return
			}
		} else {
			fmt.Println("Input yang anda berikan tidak valid")
			return
		}
	} else {
		fmt.Println("Input yang anda berikan tidak valid")
		return
	}
}
