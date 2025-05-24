package streebog

import (
	"fmt"

	"github.com/ChainsAre2Tight/streebog/pkg/constants"
	"github.com/ChainsAre2Tight/streebog/pkg/digest"
)

func Streebog256(message []byte) ([]byte, error) {
	res, err := digest.Digest(message, constants.IV256)
	if err != nil {
		return nil, fmt.Errorf("streebog.Streebog256: %s", err)
	}
	return res[32:], nil
}

func Streebog512(message []byte) ([]byte, error) {
	res, err := digest.Digest(message, constants.IV512)
	if err != nil {
		return nil, fmt.Errorf("streebog.Streebog512: %s", err)
	}
	return res, nil
}
