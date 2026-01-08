package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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