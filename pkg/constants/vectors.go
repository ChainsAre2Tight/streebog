package constants

var V512 = []uint64{
	0, 0, 0, 0, 0, 0, 0, 512,
}

var Zeroes = []uint64{
	0, 0, 0, 0, 0, 0, 0, 0,
}

var IV512 = make([]byte, 64)
var IV256 = make([]byte, 64)

func init() {
	for i := range 64 {
		IV256[i] = 1
	}
}
