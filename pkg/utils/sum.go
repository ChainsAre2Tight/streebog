package utils

import (
	"encoding/binary"
	"fmt"
	"math"
	"slices"
)

func AddInRing(dst []uint64, add []uint64) {
	if len(dst) != 8 || len(add) != 8 {
		panic(fmt.Errorf("utils.AddInRing: unexpected slice lengths. Expected: 8, Got: %d / %d", len(dst), len(add)))
	}

	buffer := make([]byte, 64)

	// TODO: Rewrite without using reverse()
	reverse := func(u []uint64) {
		for i := range 8 {
			binary.BigEndian.PutUint64(buffer[i*8:i*8+8], u[i])
		}
		slices.Reverse(buffer)
		for i := range 8 {
			u[i] = binary.BigEndian.Uint64(buffer[i*8 : i*8+8])
		}
	}

	reverse(dst)
	reverse(add)

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

	reverse(dst)
	reverse(add)
}
