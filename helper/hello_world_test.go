package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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