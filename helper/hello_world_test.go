package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"runtime"
)

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Hilmi",
			request: "Hilmi",
		},
		{
			name:    "Yahya",
			request: "Yahya",
		},
		{
			name:    "HilmiYahya",
			request: "Hilmi Yahya",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}	
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Hilmi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hilmi")
		}
	})

	b.Run("Yahya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Yahya")
		}
	})
}

func BenchmarkHelloWorldHilmi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hilmi")
	}
}

func BenchmarkHelloWorldYahya(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Yahya")
	}	
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request string
		expected string
	}{
		{
			name:     "Hilmi",
			request:  "Hilmi",
			expected: "Hello Hilmi",
		},
		{
			name:     "Yahya",
			request:  "Yahya",
			expected: "Hello Yahya",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be 'Hello Hilmi'")
		})
	}
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Hilmi", func(t *testing.T) {
		result := HelloWorld("Hilmi")
		require.Equal(t, "Hello Hilmi", result, "Result must be 'Hello Hilmi'")
	})

	t.Run("Yahya", func(t *testing.T) {
		result := HelloWorld("Yahya")
		require.Equal(t, "Hello Yahya", result, "Result must be 'Hello Yahya'")
	})
}

// Main Test / Before & After Test
func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")
	m.Run() // menjalankan semua unit test
	fmt.Println("Setelah Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows OS")
	}

	result := HelloWorld("Skip")
	if result != "Hello Skip" {
		t.Error("Result must be 'Hello Skip'")
	}
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Require'")
	fmt.Println("TestHelloWorldRequire Berjalan")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be 'Hello Assert'")
	fmt.Println("TestHelloWorldAssert Berjalan")
}

func TestHelloWorldHilmi(t *testing.T) {
	result := HelloWorld("Hilmi")
	if result != "Hello Hilmi" {
		// error
		t.Fail() // melanjutkan menjalankan kode program meskipun error, melanjutkan eksekusi unit test selanjutnya
	}
	fmt.Println("TestHelloWorldHilmi Berjalan") // kalo pale Fail maka kode ini akan dieksekusi
}

func TestHelloWorldYahya(t *testing.T) {
	result := HelloWorld("Yahya")
	if result != "Hello Yahya" {
		// error
		t.FailNow() // menghentikan menjalankan kode program jika error, melanjutkan eksekusi unit test selanjutnya
	}
	fmt.Println("TestHelloWorldYahya Berjalan") // kalo pale FailNow maka kode ini tidak akan dieksekusi
}

// Error memanggil Fail()
func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Error")
	if result != "Hello Error" {
		// error
		t.Error("Result is not Hello Error") // mencatat error namun tetap melanjutkan eksekusi
	}
	fmt.Println("TestHelloWorldError Berjalan (Dieksekusi meskipun error)")
}

// Fatal memanggil FailNow()
func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Fatal")
	if result != "Hello Fatal" {
		// error
		t.Fatal("Result is not Hello Fatal") // mencatat error dan menghentikan eksekusi
	}
	fmt.Println("TestHelloWorldFatal Berjalan (Tidak Dieksekusi jika error)")
}

// lebih baik menggunakan Error & Fatal karena bisa menambahkan pesan error (informasi kenapa unit test gagal)