package primitives

import (
	"encoding/binary"
	"fmt"
)

func P(dst []uint64) {
	if l := len(dst); l != 8 {
		panic(fmt.Sprintf("primitives.P: unexpected dst length. Expected: 8, Got: %d", l))
	}

	bytes := make([]byte, 64)
	for i := range 8 {
		binary.BigEndian.PutUint64(bytes[i*8:(i+1)*8], dst[i])
	}

	// TODO: use shared buffers allocated once per hash instance
	permuted := make([]byte, 64)
	permuted[0] = bytes[0]
	permuted[1] = bytes[8]
	permuted[2] = bytes[16]
	permuted[3] = bytes[24]
	permuted[4] = bytes[32]
	permuted[5] = bytes[40]
	permuted[6] = bytes[48]
	permuted[7] = bytes[56]
	permuted[8] = bytes[1]
	permuted[9] = bytes[9]
	permuted[10] = bytes[17]
	permuted[11] = bytes[25]
	permuted[12] = bytes[33]
	permuted[13] = bytes[41]
	permuted[14] = bytes[49]
	permuted[15] = bytes[57]
	permuted[16] = bytes[2]
	permuted[17] = bytes[10]
	permuted[18] = bytes[18]
	permuted[19] = bytes[26]
	permuted[20] = bytes[34]
	permuted[21] = bytes[42]
	permuted[22] = bytes[50]
	permuted[23] = bytes[58]
	permuted[24] = bytes[3]
	permuted[25] = bytes[11]
	permuted[26] = bytes[19]
	permuted[27] = bytes[27]
	permuted[28] = bytes[35]
	permuted[29] = bytes[43]
	permuted[30] = bytes[51]
	permuted[31] = bytes[59]
	permuted[32] = bytes[4]
	permuted[33] = bytes[12]
	permuted[34] = bytes[20]
	permuted[35] = bytes[28]
	permuted[36] = bytes[36]
	permuted[37] = bytes[44]
	permuted[38] = bytes[52]
	permuted[39] = bytes[60]
	permuted[40] = bytes[5]
	permuted[41] = bytes[13]
	permuted[42] = bytes[21]
	permuted[43] = bytes[29]
	permuted[44] = bytes[37]
	permuted[45] = bytes[45]
	permuted[46] = bytes[53]
	permuted[47] = bytes[61]
	permuted[48] = bytes[6]
	permuted[49] = bytes[14]
	permuted[50] = bytes[22]
	permuted[51] = bytes[30]
	permuted[52] = bytes[38]
	permuted[53] = bytes[46]
	permuted[54] = bytes[54]
	permuted[55] = bytes[62]
	permuted[56] = bytes[7]
	permuted[57] = bytes[15]
	permuted[58] = bytes[23]
	permuted[59] = bytes[31]
	permuted[60] = bytes[39]
	permuted[61] = bytes[47]
	permuted[62] = bytes[55]
	permuted[63] = bytes[63]

	for i := range 8 {
		dst[i] = binary.BigEndian.Uint64(permuted[i*8 : (i+1)*8])
	}
}
