package primitives

import (
	"fmt"

	"github.com/ChainsAre2Tight/streebog/pkg/tables"
)

func S(dst []uint64) {
	if len(dst) != 8 {
		panic(fmt.Errorf("primitives.S: unexpected dst length. Expected: 4, Got: %d", len(dst)))
	}
	var temp uint64
	for i := range 8 {
		temp = 0
		for l, r := 0, 56; l < 64; l, r = l+8, 48-l {
			value := (dst[i] << l) >> 56
			sub := tables.DirectSbox[byte(value)]
			temp += uint64(sub) << r
		}
		dst[i] = temp
	}
}
