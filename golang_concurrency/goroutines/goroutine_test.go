package goroutines

/*
1. Contoh sederhana goroutine
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func PrintNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func TestHelloWorld(t *testing.T) {
	PrintNumbers()
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second) //Menunggu goroutine selesai sebelum program berakhir
}
*/
/*
// 2. Goroutine dengan fungsi Anonim
func TestAnonim(t *testing.T) {
	go func() {
		fmt.Println("Goroutine dengan fungsi anonim")
	}()

	time.Sleep(1 * time.Second)
}
*/
/*
// Goroutine tanpa sinkronisasi
func Ganjil() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func Genap() {
	for i := 1; i <= 9; i += 2 {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func TestNumber(t *testing.T) {
	go Ganjil()
	go Genap()

	time.Sleep(1 * time.Second)
}
*/
/*
// Goroutine dengan parameter
func Greet(name string) {
	fmt.Println("Hello", name)
}

func TestGreet(t *testing.T) {
	go Greet("Akbar")

	time.Sleep(1 * time.Second)
}
*/
