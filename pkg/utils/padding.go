package utils

import "fmt"

func PadBytes(in []byte) []byte {
	if len(in) > 64 {
		panic(fmt.Errorf("utils.Pad: unexpected slice length: expected: <64, got: %d", len(in)))
	}
	res := make([]byte, 64)
	copy(res, in)
	if len(in) < 64 {
		res[len(in)] = 1
	}
	return res
}
