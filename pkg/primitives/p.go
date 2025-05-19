package primitives

import "github.com/ChainsAre2Tight/streebog/pkg/tables"

func P(dst []uint64) {
	var res = make([]uint64, 8)
	for i := range 8 {
		for j, index := 7, 8*i+7; j >= 0; j, index = j-1, index-1 {
			kuda := tables.PBox[index]
			main, add := kuda/8, (7-kuda%8)*8
			res[main] += uint64(byte(dst[i])) << add
			dst[i] >>= 8
		}
	}
	copy(dst, res)
}
