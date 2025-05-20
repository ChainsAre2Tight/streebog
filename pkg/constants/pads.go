package constants

var (
	IPAD []byte = make([]byte, 64)
	OPAD []byte = make([]byte, 64)
)

func init() {
	for i := range 64 {
		IPAD[i] = 0x36
		OPAD[i] = 0x5c
	}
}
