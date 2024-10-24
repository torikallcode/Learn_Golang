package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
// Table BenchMark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Muhammad",
			request: "Muhammad",
		},
		{
			name:    "Torikal",
			request: "Torikal",
		},
	}

	for _, bench := range benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.name)
			}
		})
	}
}
*/

/*
// BechMarkSub
func BenchmarkParent(b *testing.B) {
	b.Run("Muhammad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Muhammad")
		}
	})
	b.Run("Torikal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Torikal")
		}
	})
}
*/

/*
// Unit test menggunakan require
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Muhammad")
	}
}
func BenchmarkHelloWorldTorikal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Torikal")
	}
}
*/

/*
1. Menjalankan test

	func TestHelloWorld(t *testing.T) {
		result := HelloWorld("Akbar")
		if result != "Hello Akbar" {
			panic("Result is not Hello Akbar")
		}
	}
*/
/*
2. Menggagalkan test menggunakan Fail() & FailNow
func TestHelloWorldAkbar(t *testing.T) {
	result := HelloWorld("Akbar")
	if result != "Hello Akbar" {
		t.Fail()
	}
	fmt.Println("test done")
}
func TestHelloWorldTorikal(t *testing.T) {
	result := HelloWorld("Torikal")
	if result != "Hello Torikal" {
		t.FailNow()
	}
	fmt.Println("test done")
}
*/
/*
2.1 Menggagalkan test menggunakan Error() dan Fatal()
func TestHelloWorldAkbar(t *testing.T) {
	result := HelloWorld("Akbar")
	if result != "Hello Akbar" {
		t.Error("Seharusnya Hello Akbar")
	}
	fmt.Println("test done")
}
func TestHelloWorldTorikal(t *testing.T) {
	result := HelloWorld("Torikal")
	if result != "Hello Torikal" {
		t.Fatal("Seharusnya Hello Torikal")
	}
	fmt.Println("test done")
}
*/
/*
// 3. Menggunakan assert & require dari library testify
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Akbar")
	assert.Equal(t, "hai Akbar", result, "Seharusnya Hello Akbar")
	fmt.Println("test done")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Torikal")
	require.Equal(t, "hai Torikal", result, "Seharusnya Hello Torikal")
	fmt.Println("test done")
}
*/
/*
4. Membatalkan Unit test menggunakan skip()
func TestUnitSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Tidak bisa di berjalan di mac os")
	}
	result := HelloWorld("Akbar")
	assert.Equal(t, "hai Akbar", result, "Seharusnya Hallo Akbar")
}
*/
/*
5. Before & after menggunakan TestMain
func TestMain(M *testing.M) {
	fmt.Println("Sebelum Testing dimulai")
	M.Run()
	fmt.Println("Setelah Testing dimulai")
}
*/
/*
6. SubTest
func TestParent(t *testing.T) {
	t.Run("Torikal", func(t *testing.T) {
		result := HelloWorld("Torikal")
		assert.Equal(t, "hai Torikal", result, "Seharusnya Hallo Torikal")
	})
	t.Run("Akbar", func(t *testing.T) {
		result := HelloWorld("Akbar")
		assert.Equal(t, "hai Akbar", result, "Seharusnya Hallo Akbar")
	})
}
*/

// 7. Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		reuest   string
		expected string
	}{
		{
			name:     "Muhammad",
			reuest:   "Muhammad",
			expected: "Hallo Muhammad",
		}, {
			name:     "Torikal",
			reuest:   "Torikal",
			expected: "Hallo Torikal",
		},
		{
			name:     "Akbar",
			reuest:   "Akbar",
			expected: "Hallo Akbar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.reuest)
			assert.Equal(t, test.expected, result)
		})
	}
}
