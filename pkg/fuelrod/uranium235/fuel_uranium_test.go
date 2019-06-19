package uranium235

import (
	"testing"
)

func BenchmarkFission(b *testing.B) {
	u := Uranium235Oxide{
		Mass:    52000,
		Quality: 3,
	}
	for n := 0; n < b.N; n++ {
		u.DoFission()
	}
}
