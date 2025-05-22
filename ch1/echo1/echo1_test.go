package main

import (
	"os"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	os.Args = []string{"cmd", "test-value", "another string"}
	for i := 0; i < b.N; i++ {
		echo1()
	}
}
