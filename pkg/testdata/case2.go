package testdata

import "encoding/hex"

var M2 = []byte{}

func init() {
	M2, _ = hex.DecodeString("fbe2e5f0eee3c820fbeafaebef20fffbf0e1e0f0f520e0ed20e8ece0ebe5f0f2f120fff0eeec20f120faf2fee5e2202ce8f6f3ede220e8e6eee1e8f0f2d1202ce8f0f2e5e220e5d1")
}

var M2_chunk_1 = []uint64{
	0xfbeafaebef20fffb, 0xf0e1e0f0f520e0ed,
	0x20e8ece0ebe5f0f2, 0xf120fff0eeec20f1,
	0x20faf2fee5e2202c, 0xe8f6f3ede220e8e6,
	0xeee1e8f0f2d1202c, 0xe8f0f2e5e220e5d1,
}

var Case2_EPSILON_1 = []uint64{
	0xfbeafaebef20fffb, 0xf0e1e0f0f520e0ed,
	0x20e8ece0ebe5f0f2, 0xf120fff0eeec20f1,
	0x20faf2fee5e2202c, 0xe8f6f3ede220e8e6,
	0xeee1e8f0f2d1202c, 0xe8f0f2e5e220e5d1,
}

var M2_chunk_2 = []uint64{
	0x0000000000000000, 0x0000000000000000,
	0x0000000000000000, 0x0000000000000000,
	0x0000000000000000, 0x0000000000000000,
	0x0000000000000001, 0xfbe2e5f0eee3c820,
}

var Case2_EPSILON_2 = []uint64{
	0xfbeafaebef20fffb, 0xf0e1e0f0f520e0ed,
	0x20e8ece0ebe5f0f2, 0xf120fff0eeec20f1,
	0x20faf2fee5e2202c, 0xe8f6f3ede220e8e6,
	0xeee1e8f0f2d1202e, 0xe4d3d8d6d104adf1,
}

var Case2_HASH_512 = []byte{}
var Case2_HASH_256 = []byte{}

func init() {
	Case2_HASH_512 = d("28fbc9bada033b1460642bdcddb90c3fb3e56c497ccd0f62b8a2ad4935e85f037613966de4ee00531ae60f3b5a47f8dae06915d5f2f194996fcabf2622e6881e")
	Case2_HASH_256 = d("508f7e553c06501d749a66fc28c6cac0b005746d97537fa85d9e40904efed29d")
}
