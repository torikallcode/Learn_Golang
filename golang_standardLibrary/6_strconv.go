package main

func main() {

	// result, err := strconv.ParseBool("true")
	// if err != nil {
	// 	fmt.Println("Error", err.Error())
	// } else {
	// 	fmt.Println(result)
	// }

	// resultInt, err := strconv.Atoi("9999")
	// if err != nil {
	// 	fmt.Println("Error", err.Error())
	// } else {
	// 	fmt.Println(resultInt)
	// }

	// binary := strconv.FormatInt(9999, 2)
	// fmt.Println(binary)

	// var angka int
	// fmt.Print("Masukkan sebuah angka: ")
	// fmt.Scan(&angka)

	// result := strconv.Itoa(angka)
	// fmt.Println(result)
	// fmt.Println(reflect.TypeOf(result))

}

/*
bagaimana cara kita melakukan konversi yang tipe datanya berbeda? misal dari int ke string
atau sebaliknya , hal tersebut bisa kita lakukan dengan bantuan package strconv(string conversion)
yang sering digunakan yaitu
i, err := strconv.Atoi("-42") //string to int
s := strconv.Itoa(-42) // int to string
b, err := strconv.ParseBool("true") //string to boolean
f, err := strconv.ParseFloat("3.1415", 64) //string to float64
i, err := strconv.ParseInt("-42", 10, 64) // string to int
s := strconv.FormatBool(true) // bool ke string
s := strconv.FormatFloat(3.1415, 'E', -1, 64) // float ke string
s := strconv.FormatInt(-42, 16) //int ke string
// contoh 1
	result, err := strconv.ParseBool("salah")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}
	output = true
// contoh 2
	resultInt, err := strconv.Atoi("9999")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}
output 9999
// contoh 3
	binary := strconv.FormatInt(9999, 2)
	fmt.Println(binary)
	output = 10011100001111
*/
