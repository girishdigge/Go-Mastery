package iteration

import (
	"strings"
	"testing"
)

const repeatCount = 5

func Repeat(charcter string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(charcter)
	}
	return repeated.String()
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
