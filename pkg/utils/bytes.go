package utils

import (
	"encoding/binary"
	"fmt"
)

func BytesToUints(in []byte, dst []uint64) {
	if len(in) != 64 {
		panic(fmt.Errorf("utils.BytesToUints: unexpected in slice length. Expected: 64, Got: %d", len(in)))
	}
	if len(dst) != 8 {
		panic(fmt.Errorf("utils.BytesToUints: unexpected dst slice length. Expected: 8, Got: %d", len(in)))
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

func UintsToBytes(in []uint64) []byte {
	if len(in) != 8 {
		panic(fmt.Errorf("utils.UintsToBytes: unexpected slice length. Expected: 8, Got: %d", len(in)))
	}
	res := make([]byte, 0, 64)
	for i := range 8 {
		res = binary.BigEndian.AppendUint64(res, in[i])
	}
	return res
}
