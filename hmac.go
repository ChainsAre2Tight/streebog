package streebog

import (
	"fmt"

	"github.com/ChainsAre2Tight/streebog/pkg/constants"
	"github.com/ChainsAre2Tight/streebog/pkg/utils"
)

func HMAC256(K, T []byte) ([]byte, error) {
	hmac, err := hmac_common(K, T, Streebog256)
	if err != nil {
		return nil, fmt.Errorf("streebog.HMAC256: %s", err)
	}
	return hmac, nil
}

func HMAC512(K, T []byte) ([]byte, error) {
	hmac, err := hmac_common(K, T, Streebog512)
	if err != nil {
		return nil, fmt.Errorf("streebog.HMAC512: %s", err)
	}
	return hmac, nil
}

func hmac_common(K, T []byte, hashFunc func(message []byte) ([]byte, error)) ([]byte, error) {

	fail := func(err error) ([]byte, error) {
		return nil, fmt.Errorf("streebog.hmac_common: %s", err)
	}

	k_ipad := utils.XORBytes(K, constants.IPAD)
	k_ipad = append(k_ipad, T...)
	hash, err := hashFunc(k_ipad)
	if err != nil {
		return fail(fmt.Errorf("hash(K ^ ipad | T): %s", err))
	}

	k_opad := utils.XORBytes(K, constants.OPAD)
	k_opad = append(k_opad, hash...)
	hash2, err := hashFunc(k_opad)
	if err != nil {
		return fail(fmt.Errorf("hash(K ^ opad | hash(K ^ ipad | T)): %s", err))
	}
	return hash2, nil
}
