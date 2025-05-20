package testdata

import (
	"encoding/hex"
	"fmt"
)

func d(in string) []byte {
	res, err := hex.DecodeString(in)
	if err != nil {
		panic(fmt.Errorf("streebog_test.d: %s", err))
	}
	return res
}
