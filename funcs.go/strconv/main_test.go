package main

import (
	"fmt"
	"strconv"
	"testing"
)

// go test -bench=. -benchtime=1s -benchmem
func BenchmarkSprintf(b *testing.B) {
	val := 66.00
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%.2f", val)
	}
}

func BenchmarkStrconv(b *testing.B) {
	val := 66.00
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatFloat(val, 'f', 2, 64)
	}
}
