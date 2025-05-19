package primitives

import "fmt"

func X(dst, target []uint64) {
	if len(dst) != 8 || len(target) != 8 {
		panic(fmt.Errorf("primitives.X: unexpected slice lengths: %d, %d expected: 8", len(dst), len(target)))
	}

	for i := range 8 {
		dst[i] ^= target[i]
	}
}
