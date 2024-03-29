package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
~ Running semua Benchmark dan Unit Test 			>>  go test -v -bench=.
~ Running semua Benchmarknya saja 					>>  go test -v -run=UnitTestYgTidakAda -bench=.
~ Running satu Benchmark tertentu 					>>  go test -v -run=UnitTestYgTidakAda -bench=NamaBenchmarkTest
~ Running semua Benchmark di root / semua modul		>>  go test -v -run=UnitTestYgTidakAda -bench=. ./...
*/
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fadli")
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Fadli", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
            HelloWorld("Fadli")
        }
	})

	b.Run("Darusalam", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
            HelloWorld("Darusalam")
        }
	})
}

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "Fadli",
			request: "Fadli",
		},
		{
			name: "Fadli",
			request: "Fadli",
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

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fadli",
			request:  "Fadli",
			expected: "Hello Fadli",
		},
		{
			name:     "Darusalam",
			request:  "Darusalam",
			expected: "Hello Darusalam",
		},
		{
			name:     "Sragen",
			request:  "Sragen",
			expected: "Hello Sragen",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Untuk run satu subtest saja : go test -v -run=TestSubTest/fadli
// Untuk run semua subtest dengan nama fadli : go test -v -run=/fadli
func TestSubTest(t *testing.T) {
	t.Run("Fadli", func(t *testing.T) {
		result := HelloWorld("Fadli")
		require.Equal(t, "Hello Fadli", result, "Result should be Hello Fadli")
	})

	t.Run("Darusalam", func(t *testing.T) {
		result := HelloWorld("Darusalam")
		require.Equal(t, "Hello Darusalam", result, "Result should be Hello Darusalam")
	})

}

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

func Iterative(n int) {
	for i := n; i >= 0; i-- {
        // fmt.Println(i)
    }
}

func Recursive(n int) {
	if n >= 0 {
		// fmt.Println(n)
        Recursive(n - 1)
    }
}

func BenchmarkIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Iterative(20)
    }
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Recursive(20)
    }
}