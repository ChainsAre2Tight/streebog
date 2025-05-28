package testdata

import (
	"encoding/hex"
	"slices"
)

var M2 = []byte{}

func init() {
	M2, _ = hex.DecodeString("fbe2e5f0eee3c820fbeafaebef20fffbf0e1e0f0f520e0ed20e8ece0ebe5f0f2f120fff0eeec20f120faf2fee5e2202ce8f6f3ede220e8e6eee1e8f0f2d1202ce8f0f2e5e220e5d1")
	slices.Reverse(M2)
}

var M2_chunk_1 = []uint64{
	0xd1e520e2e5f2f0e8, 0x2c20d1f2f0e8e1ee, 0xe6e820e2edf3f6e8, 0x2c20e2e5fef2fa20,
	0xf120eceef0ff20f1, 0xf2f0e5ebe0ece820, 0xede020f5f0e0e1f0, 0xfbff20efebfaeafb,
}

var Case2_EPSILON_1 = []uint64{
	0xd1e520e2e5f2f0e8, 0x2c20d1f2f0e8e1ee, 0xe6e820e2edf3f6e8, 0x2c20e2e5fef2fa20,
	0xf120eceef0ff20f1, 0xf2f0e5ebe0ece820, 0xede020f5f0e0e1f0, 0xfbff20efebfaeafb,
}

var M2_chunk_2 = []uint64{
	0x20c8e3eef0e5e2fb, 0x0100000000000000, 0x0000000000000000, 0x0000000000000000,
	0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000,
}

var Case2_EPSILON_2 = []uint64{
	0xf1ad04d1d6d8d3e4, 0x2e20d1f2f0e8e1ee, 0xe6e820e2edf3f6e8, 0x2c20e2e5fef2fa20,
	0xf120eceef0ff20f1, 0xf2f0e5ebe0ece820, 0xede020f5f0e0e1f0, 0xfbff20efebfaeafb,
}

var Case2_HASH_512 = []byte{}
var Case2_HASH_256 = []byte{}

func init() {
	Case2_HASH_512 = d("28fbc9bada033b1460642bdcddb90c3fb3e56c497ccd0f62b8a2ad4935e85f037613966de4ee00531ae60f3b5a47f8dae06915d5f2f194996fcabf2622e6881e")
	Case2_HASH_256 = d("508f7e553c06501d749a66fc28c6cac0b005746d97537fa85d9e40904efed29d")
	slices.Reverse(Case2_HASH_256)
	slices.Reverse(Case2_HASH_512)
}
