package primitives

import "github.com/ChainsAre2Tight/streebog/pkg/tables"

func L(dst []uint64) {
	for i := range 8 {
		dst[i] = linear(dst[i])
	}
}

func linear(in uint64) uint64 {
	var res uint64 = 0
	index := 63
	for in > 0 {
		if in&1 > 0 {
			res ^= tables.Linear[index]
		}
		in >>= 1
		index--
	}
	return res
}
