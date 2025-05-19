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
	primitives.X(h, N)
	primitives.S(h)
	primitives.P(h)
	primitives.L(h)

	// k1
	buffer_K := make([]uint64, 8)
	copy(buffer_K, h)

	primitives.X(h, m)
	primitives.S(h)
	primitives.P(h)
	primitives.L(h)

	for round := range 11 {
		// calc K_i
		primitives.X(buffer_K, constants.All[round])
		primitives.S(buffer_K)
		primitives.P(buffer_K)
		primitives.L(buffer_K)

		primitives.X(h, buffer_K)
		primitives.S(h)
		primitives.P(h)
		primitives.L(h)
	}

	// k13
	primitives.X(buffer_K, constants.All[11])
	primitives.S(buffer_K)
	primitives.P(buffer_K)
	primitives.L(buffer_K)

	primitives.X(h, buffer_K)

	primitives.X(h, h_temp)
	primitives.X(h, m)
}
