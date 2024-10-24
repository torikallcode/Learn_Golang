package channels

/*
// Contoh sederhana
func TestBuffered(t *testing.T) {
	ch := make(chan int, 3) //channel dengan kapasitas 3

	// mengirim nilai ke channel
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("Buffered data sent")

	// menerima nilai dari channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
// Blocking pada buffered channels
func TestBuffered(t *testing.T) {
	ch := make(chan int, 2) // channel dengan kapasitas 2

	// mengirim nilai ke chanel
	ch <- 1
	ch <- 2
	fmt.Println("Channel penuh , akan menunggu penerima") //chanel ngeblock karena kapasitas penuh
	fmt.Println("Data yang sudah diterima:", <-ch) // menerima channel agar channel bisa menerima nilai karena kapasitasnya sudah lega

	// mengirim nilai baru
	ch <- 3
	fmt.Println("Data ketiga berhasil dikirim")
}

		// Buffer penuh
    ch := make(chan int, 2)  // Buffer dengan kapasitas 2

    ch <- 1
    ch <- 2

    fmt.Println("Buffer penuh. Akan menunggu...")

    // Pengiriman ini akan blocking karena buffer sudah penuh
    ch <- 3  // Menunggu hingga ada yang menerima data
    fmt.Println("Data ketiga berhasil dikirim.")

		// Buffer kosong
		    ch := make(chan int, 2)  // Buffer dengan kapasitas 2

    ch <- 1
    ch <- 2

    fmt.Println(<-ch)  // Mengambil nilai pertama dari buffer (1)
    fmt.Println(<-ch)  // Mengambil nilai kedua dari buffer (2)

    fmt.Println("Buffer kosong. Akan menunggu...")

    // Ini akan blocking sampai ada data baru yang dikirim ke channel
    fmt.Println(<-ch)  // Menunggu data masuk
*/

/*
// Buffered channels dengan goroutine
func SendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Mengirim:", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}
	defer close(ch)
}

func TestBuffered(t *testing.T) {
	ch := make(chan int, 2)

	go SendData(ch)
	for data := range ch {
		fmt.Println("Menerima data:", data)
		time.Sleep(2 * time.Second)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Selesai menerima semua data")
}
*/
