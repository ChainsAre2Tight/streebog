package utils

import "fmt"

func PadBytes(in []byte, dst []byte) []byte {
	if l := len(in); l > 64 {
		panic(fmt.Errorf("utils.Pad: unexpected in slice length: expected: <64, got: %d", l))
	}
	if l := len(dst); l != 64 {
		panic(fmt.Errorf("utils.Pad: unexpected dst slice length: expected: 64, got: %d", l))
	}

	for i := range dst {
		dst[i] = 0
	}

	copy(dst, in)
	if len(in) < 64 {
		dst[len(in)] = 1
	}
	return dst
}
