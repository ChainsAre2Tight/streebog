package utils

import (
	"encoding/binary"
	"fmt"
)

func BytesToUints(in []byte) ([]uint64, error) {
	if len(in) != 64 {
		return nil, fmt.Errorf("utils.BytesToUints: unexpected slice length. Expected: 64, Got: %d", len(in))
	}
	res := make([]uint64, 8)
	for i := range 8 {
		res[i] = binary.BigEndian.Uint64(in[8*i : 8*i+8])
	}

	return res, nil
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
