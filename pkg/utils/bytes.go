package utils

import (
	"encoding/binary"
	"fmt"
)

func BytesToUints(in []byte, dst []uint64) {
	if l := len(in); l != 64 {
		panic(fmt.Errorf("utils.BytesToUints: unexpected in slice length. Expected: 64, Got: %d", l))
	}
	if l := len(dst); l != 8 {
		panic(fmt.Errorf("utils.BytesToUints: unexpected dst slice length. Expected: 8, Got: %d", l))
	}
	dst[0] = binary.BigEndian.Uint64(in[0:8])
	dst[1] = binary.BigEndian.Uint64(in[8:16])
	dst[2] = binary.BigEndian.Uint64(in[16:24])
	dst[3] = binary.BigEndian.Uint64(in[24:32])
	dst[4] = binary.BigEndian.Uint64(in[32:40])
	dst[5] = binary.BigEndian.Uint64(in[40:48])
	dst[6] = binary.BigEndian.Uint64(in[48:56])
	dst[7] = binary.BigEndian.Uint64(in[56:])
}

func UintsToBytes(in []uint64, dst []byte) {
	if l := len(in); l != 8 {
		panic(fmt.Errorf("utils.UintsToBytes: unexpected in slice length. Expected: 8, Got: %d", l))
	}
	if l := len(dst); l != 64 {
		panic(fmt.Errorf("utils.UintsToBytes: unexpected dst slice length. Expected: 64, Got: %d", l))
	}

	binary.BigEndian.PutUint64(dst[:8], in[0])
	binary.BigEndian.PutUint64(dst[8:16], in[1])
	binary.BigEndian.PutUint64(dst[16:24], in[2])
	binary.BigEndian.PutUint64(dst[24:32], in[3])

	binary.BigEndian.PutUint64(dst[32:40], in[4])
	binary.BigEndian.PutUint64(dst[40:48], in[5])
	binary.BigEndian.PutUint64(dst[48:56], in[6])
	binary.BigEndian.PutUint64(dst[56:64], in[7])
}
