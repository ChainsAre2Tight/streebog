package round

import (
	"github.com/ChainsAre2Tight/streebog/pkg/constants"
	"github.com/ChainsAre2Tight/streebog/pkg/primitives"
)

// writes all to h
func G(h, m, N []uint64) {
	h_temp := make([]uint64, 8)
	copy(h_temp, h)

	// pre iterations
	xspl(h, N)

	// k1
	buffer_K := make([]uint64, 8)
	copy(buffer_K, h)

	xspl(h, m)

	for round := range 11 {
		xspl(buffer_K, constants.All[round])
		xspl(h, buffer_K)
	}

	// k13
	xspl(buffer_K, constants.All[11])
	primitives.X(h, buffer_K)

	primitives.X(h, h_temp)
	primitives.X(h, m)
}

func xspl(buffer, constant []uint64) {
	primitives.X(buffer, constant)
	primitives.S(buffer)
	primitives.P(buffer)
	primitives.L(buffer)
}
