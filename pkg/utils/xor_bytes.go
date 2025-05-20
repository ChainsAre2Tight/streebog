package utils

func XORBytes(v1, v2 []byte) []byte {
	if len(v1) > len(v2) {
		res := make([]byte, len(v1))
		i := 0
		for ; i < len(v2); i++ {
			res[i] = v1[i] ^ v2[i]
		}
		for ; i < len(v1); i++ {
			res[i] = v1[i]
		}
		return res
	} else {
		res := make([]byte, len(v2))
		i := 0
		for ; i < len(v1); i++ {
			res[i] = v1[i] ^ v2[i]
		}
		for ; i < len(v2); i++ {
			res[i] = v2[i]
		}
		return res
	}
}
