package utils

import "fmt"

func BytesToUints(in []byte) ([]uint64, error) {
	if len(in) != 64 {
		return nil, fmt.Errorf("utils.BytesToUints: unexpected slice length. Expected: 64, Got: %d", len(in))
	}
	res := make([]uint64, 8)
	for i := range 8 {
		for j, k := i*8, 56; j < i*8+8; j, k = j+1, k-8 {
			res[i] += uint64(in[j]) << k
		}
	}
	return res, nil
}

func UintsToBytes(in []uint64) []byte {
	if len(in) != 8 {
		panic(fmt.Errorf("utils.UintsToBytes: unexpected slice length. Expected: 8, Got: %d", len(in)))
	}
	res := make([]byte, 64)
	for i := range 8 {
		buffer := in[i]
		i8 := 8 * i
		for j := 7; j >= 0; j-- {
			res[i8+j] = byte(buffer)
			buffer >>= 8
		}
	}
	return res
}
