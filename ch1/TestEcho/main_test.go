// Функции производительности различных реализаций Echo
package TestEcho

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkEcho1(b *testing.B) {
	start := time.Now()
	for i := 1; i < b.N; i++ {
		echo1()
	}
	end := time.Since(start).Seconds()
	fmt.Println("echo1 time: ", end)
}

func BenchmarkEcho2(b *testing.B) {
	start := time.Now()
	for i := 1; i < b.N; i++ {
		echo2()
	}
	end := time.Since(start).Seconds()
	fmt.Println("echo2 time: ", end)
}

func BenchmarkEcho3(b *testing.B) {
	start := time.Now()
	for i := 1; i < b.N; i++ {
		echo3()
	}
	end := time.Since(start).Seconds()
	fmt.Println("echo3 time: ", end)
}

// results
// echo v1: 28560             56158 ns/op				1.1339227809999999s
// echo v2: 32188             42502 ns/op				1.229379511s
// echo v3: 29407             39424 ns/op				1.055014898s
