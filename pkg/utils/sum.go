package utils

import (
	"fmt"
	"math"
)

// TODO: write tests and debug....
func AddInRing(dst []uint64, add []uint64) {
	if len(dst) != 8 || len(add) != 8 {
		panic(fmt.Errorf("utils.AddInRing: unexpected slice lengths. Expected: 8, Got: %d / %d", len(dst), len(add)))
	}
	var carry uint64 = 0
	for i := 7; i >= 0; i-- {
		if add[i] > math.MaxUint64-carry-dst[i] {
			dst[i] += add[i] + carry
			carry = 1
		} else {
			dst[i] += add[i] + carry
			carry = 0
		}

	}
}
