package primitives

import "github.com/ChainsAre2Tight/streebog/pkg/tables"

func L(dst []uint64) {
	for i := range 8 {
		dst[i] = linear(dst[i])
	}
}

func linear(in uint64) uint64 {
	return tables.LinearLookup[0][byte(in>>56)] ^
		tables.LinearLookup[1][byte(in>>48)] ^
		tables.LinearLookup[2][byte(in>>40)] ^
		tables.LinearLookup[3][byte(in>>32)] ^
		tables.LinearLookup[4][byte(in>>24)] ^
		tables.LinearLookup[5][byte(in>>16)] ^
		tables.LinearLookup[6][byte(in>>8)] ^
		tables.LinearLookup[7][byte(in)]
}
