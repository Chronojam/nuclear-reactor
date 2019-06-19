package uranium235

import (
	"testing"
)

func TestU235(t *testing.T) {
	u := Uranium235Oxide{
		Mass:    52000,
		Quality: 3,
	}
	for i := 0; i < 100000000; i++ {
		u.DoFission()
	}
	t.Logf("%v", u.DoFission())
	t.Logf("%v", u.Quality)
}

func BenchmarkFission(b *testing.B) {
	u := Uranium235Oxide{
		Mass:    52000,
		Quality: 3,
	}
	for n := 0; n < b.N; n++ {
		u.DoFission()
	}
}
