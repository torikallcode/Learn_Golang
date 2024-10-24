package main

func kuadrat(n []int) int {
	hasil := 0
	for _, angka := range n {
		hasil += angka * angka
	}
	return hasil
}

func main() {
	input := []int{1, 2, 2}
	print(kuadrat(input))
}
