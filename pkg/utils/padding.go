package utils

import "fmt"

func PadBytes(in []byte) []byte {
	if len(in) > 64 {
		panic(fmt.Errorf("utils.Pad: unexpected slice length: expected: <64, got: %d", len(in)))
	}
	res := make([]byte, 64)
	var c = 63
	for i := len(in) - 1; i >= 0; i-- {
		res[c] = in[i]
		c--
	}
	res[c] = 1
	return res
}
