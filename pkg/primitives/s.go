package primitives

import (
	"fmt"

	"github.com/ChainsAre2Tight/streebog/pkg/tables"
)

func S(dst []uint64) {
	if len(dst) != 8 {
		panic(fmt.Errorf("primitives.S: unexpected dst length. Expected: 8, Got: %d", len(dst)))
	}

	for i := range 8 {
		var temp uint64
		for byteIndex := 0; byteIndex < 8; byteIndex++ {
			shift := (7 - byteIndex) * 8
			b := byte(dst[i] >> shift)
			sub := tables.DirectSbox[b]
			temp |= uint64(sub) << shift
		}
		dst[i] = temp
	}
}
