package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Jika terdapat function TestMain, maka secara otomatis Go-Lang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package
// TestMain dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test
func TestMain(m *testing.M) {
	// Before Testing
	fmt.Println("Before Unit Test")

	m.Run()

	// After Testing
	fmt.Println("After Unit Test")
}

// Jika OS nya adalah Mac OS, maka test-nya di skip / dilewati
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" { // GO Operation System mac os = darwin
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Fadli")
	require.Equal(t, "Hello Fadli", result, "Result should be Hello Fadli")
}

// Require menggunakan failNow() / jika error maka kode tidak dilanjutkan
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fadli")
	require.Equal(t, "Hello Fadli", result, "Result should be Hello Fadli")
	fmt.Println("TestHelloWorld with Assert Done")
}

// Assert menggunakan fail() / jika error maka kode tetap dilanjutkan
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fadli")
	assert.Equal(t, "Hello Fadli", result, "Result should be Hello Fadli")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldFadli(t *testing.T) {
	result := HelloWorld("Fadli")

	if result != "Hello Fadli" {
		// error
		t.Error("Result should be Hello Fadli")
	}
}

func TestHelloWorldDarusalam(t *testing.T) {
	result := HelloWorld("Darusalam")

	if result != "Hello Darusalam" {
		// error
		t.Fatal("Result should be Hello Darusalam")
	}
}
