package primitives

import (
	"fmt"

	"github.com/ChainsAre2Tight/streebog/pkg/tables"
)

func S(dst []uint64) {
	if l := len(dst); l != 8 {
		panic(fmt.Errorf("primitives.S: unexpected dst length. Expected: 8, Got: %d", l))
	}

	for i := range 8 {
		var temp uint64 = dst[i]
		var subbed uint64 = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		subbed = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		subbed = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		subbed = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		subbed = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		subbed = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		subbed = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		subbed = tables.DirectSbox[temp&0xff]
		temp >>= 8
		temp |= subbed
		dst[i] = temp
	}
}
