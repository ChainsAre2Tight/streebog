package primitives

import "github.com/ChainsAre2Tight/streebog/pkg/tables"

func P(dst []uint64) {
	var res [8]uint64

	for i := range 64 {
		srcBit := i
		dstBit := tables.PBox[i]

		srcWord := srcBit / 8
		srcShift := (7 - (srcBit % 8)) * 8

		dstWord := dstBit / 8
		dstShift := (7 - (dstBit % 8)) * 8

		b := byte(dst[srcWord] >> srcShift)
		res[dstWord] |= uint64(b) << dstShift
	}

	copy(dst, res[:])
}
