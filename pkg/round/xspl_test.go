package round

import "testing"

func BenchmarkXSPL(b *testing.B) {
	in := make([]uint64, 8)
	add := make([]uint64, 8)
	for b.Loop() {
		xspl(in, add)
	}
}
