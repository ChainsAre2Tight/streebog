package primitives

import (
	"encoding/binary"
	"fmt"

	"github.com/ChainsAre2Tight/streebog/pkg/tables"
)

func P(dst []uint64) {
	if l := len(dst); l != 8 {
		panic(fmt.Sprintf("primitives.P: unexpected dst length. Expected: 8, Got: %d", l))
	}

	bytes := make([]byte, 64)
	for i := range 8 {
		binary.BigEndian.PutUint64(bytes[i*8:(i+1)*8], dst[i])
	}

	permuted := make([]byte, 64)
	for i := range 64 {
		permuted[i] = bytes[tables.PBox[i]]
	}

	for i := range 8 {
		dst[i] = binary.BigEndian.Uint64(permuted[i*8 : (i+1)*8])
	}
}
