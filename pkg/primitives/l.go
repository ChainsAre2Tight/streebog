package primitives

import (
	"math/bits"

	"github.com/ChainsAre2Tight/streebog/pkg/tables"
)

func L(dst []uint64) {
	for i := range 8 {
		dst[i] = linear(dst[i])
	}
}

func linear(in uint64) uint64 {
	var res uint64
	for in != 0 {
		idx := bits.LeadingZeros64(in)
		res ^= tables.Linear[idx]
		in &^= 1 << (63 - idx)
	}
	return res
}
