package primitives

import (
	"fmt"

	"github.com/viant/vec/bitwise"
)

func X(dst, target []uint64) {
	if len(dst) != 8 || len(target) != 8 {
		panic(fmt.Errorf("primitives.X: unexpected slice lengths: %d, %d expected: 8", len(dst), len(target)))
	}

	bitwise.Uint64s(dst).XorAVX2(dst, target)
}
