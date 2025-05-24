package digest

import (
	"fmt"
	"slices"

	"github.com/ChainsAre2Tight/streebog/pkg/constants"
	"github.com/ChainsAre2Tight/streebog/pkg/round"
	"github.com/ChainsAre2Tight/streebog/pkg/utils"
)

func Digest(message []byte, iv []byte) ([]byte, error) {
	M := message
	fail := func(err error) ([]byte, error) {
		return nil, fmt.Errorf("digest.Digest: %s", err)
	}
	h, err := utils.BytesToUints(iv)
	if err != nil {
		return fail(fmt.Errorf("iv: %s", err))
	}
	N := make([]uint64, 8)
	summ := make([]uint64, 8)

	slices.Reverse(message)

	// 2.1
	for len(M) >= 64 {
		// 2.2
		l := len(M) - 64
		var temp []byte
		temp, M = M[l:], M[:l]
		m, err := utils.BytesToUints(temp)
		if err != nil {
			return fail(fmt.Errorf("digestion at length %d: %s", l, err))
		}
		// 2.3
		round.G(h, m, N)

		// 2.4
		utils.AddInRing(N, constants.V512)

		// 2,5
		utils.AddInRing(summ, m) // error begins here
	}

	// 3.1
	m, err := utils.BytesToUints(utils.PadBytes(M))
	if err != nil {
		return fail(fmt.Errorf("step 3.1: %s", err))
	}

	// 3.2
	round.G(h, m, N)

	// 3.3
	utils.AddInRing(N, []uint64{0, 0, 0, 0, 0, 0, 0, uint64(len(M) * 8)}) // and here

	// 3.4
	utils.AddInRing(summ, m) // off by one in the last bit

	// 3.5
	round.G(h, N, constants.Zeroes)

	// 3.6
	round.G(h, summ, constants.Zeroes)

	b := utils.UintsToBytes(h)
	slices.Reverse(b)

	return b, nil
}
