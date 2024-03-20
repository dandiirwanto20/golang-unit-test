package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// implement table benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Dandi",
			request: "Dandi",
		},
		{
			name:    "Windah",
			request: "Windah",
		},
		{
			name:    "Irwanto",
			request: "Irwanto",
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

// implement sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Dandi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dandi")
		}
	})
	b.Run("Irwanto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Irwanto")
		}
	})
}

// menjalankan benchmark dan unit test: go test -v -bench=.
// menjalankan benchmark saja: go test -v -run=NotMacthUnitTest -bench=.
// menjalankan specific benchmark: go test -v -run=NotMacthUnitTest -bench=BenchmarkHelloWorld
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dandi")
	}
}

func BenchmarkHelloWorldDandi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("World")
	}
}

// implemnt mock unit test for case difficult object test like (API call) or connection databases

// implement table test concept alternative for redudan implement unit test manually
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Dandi",
			request:  "Dandi",
			expected: "Hello Dandi",
		},
		{
			name:     "Irwanto",
			request:  "Irwanto",
			expected: "Hello Irwanto",
		},
		{
			name:     "Windah",
			request:  "Windah",
			expected: "Hello Windah",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// implement Sub Test unit test

// jika ingin menjalankan hanya salah satu subtestnya gunakan: go test -v -run=TestNamaFuncton/NamaSubtest
// atau semua test subtest di semua func: go test -v -run=/NamaSubTest
func TestSubTest(t *testing.T) {
	// Dandi is name of subtest
	t.Run("Dandi", func(t *testing.T) {
		result := HelloWorld("Dandi")
		require.Equal(t, "Hello Dandi", result, "Result must be Hello World")
	})
	t.Run("Irwanto", func(t *testing.T) {
		result := HelloWorld("Irwanto")
		require.Equal(t, "Hello Irwanto", result, "Result must be Hello World")
	})
}

// for make before after unit test
func TestMain(m *testing.M) {
	// before unit test
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

// untuk skip sebuah unit test jika tidak diperlukan
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows/amd64" {
		t.Skip("Can not run in Windows")
	}

	result := HelloWorld("World")
	require.Equal(t, "Hello World", result, "Result must be Hello World")
}

// implement Assertion unit test alternatif for selection (if else) and return value is Fail()
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("World")
	assert.Equal(t, "Hello World", result, "Result must be Hello World")
	fmt.Println("TestHelloWorld with Assert Done")
}

// implement Require unit test alternative for selection (if else) and return value is FailNow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("World")
	require.Equal(t, "Hello World", result, "Result must be Hello World")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	// result := HelloWorld("Dandi")

	if result != "Hello World" {
		// implement error for Fail and can continue at the program
		// t.Fail()

		// implement Error include Fail()
		t.Error("Result must be Hello World")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloDandi(t *testing.T) {
	result := HelloWorld("Dandi")
	// result := HelloWorld("World")

	if result != "Hello Dandi" {
		// implement error for FailNow() langsung di exec now
		// t.FailNow()

		// implement Fatal include FailNow()
		t.Fatal("Result must be Hello Dandi")
	}
}

// command for test:
// for all unit test -> go test,

// show function in unit test -> go test -v (v is verbose), and

// go test -v -run [name of unit test function] or go test -v -run=[name of unit test function] (can show same prefix) -> for specific func

// command for root folder:
// go test ./...
