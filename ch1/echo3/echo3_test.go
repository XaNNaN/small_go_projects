package main

import (
	"os"
	"testing"
)

func BenchmarkEcho3(b *testing.B) {
	os.Args = []string{"cmd", "test-value", "another string"}
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
