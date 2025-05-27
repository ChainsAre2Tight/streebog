package streebog

import (
	"encoding/binary"
	"hash"
	"slices"

	"github.com/ChainsAre2Tight/streebog/pkg/constants"
	"github.com/ChainsAre2Tight/streebog/pkg/round"
	"github.com/ChainsAre2Tight/streebog/pkg/utils"
)

var _ hash.Hash = (*Streebog)(nil)

type Streebog struct {
	size int

	// used to reverse message in current implementation
	bufferMessage []byte

	// buffers for Write()
	bufferH    []uint64
	bufferN    []uint64
	bufferSumm []uint64
	bufferM    []uint64

	// buffers for Sum()
	bufferOutH    []uint64
	bufferOutN    []uint64
	bufferOutSumm []uint64
}

func (h *Streebog) BlockSize() int {
	return 64
}

func (h *Streebog) Size() int {
	return h.size
}
func New(size int) *Streebog {
	if size != 32 && size != 64 {
		panic("invalid hash size")
	}
	res := &Streebog{
		size: size,
	}
	res.Reset()
	return res
}

func (h *Streebog) Write(p []byte) (n int, err error) {
	h.bufferMessage = make([]byte, len(p))
	copy(h.bufferMessage, p)

	// 2.1
	var c = 0
	for len(h.bufferMessage) >= 64 {
		c += 64
		// 2.2
		var temp []byte
		temp, h.bufferMessage = h.bufferMessage[:64], h.bufferMessage[64:]
		utils.BytesToUints(temp, h.bufferM)

		// 2.3
		round.G(h.bufferH, h.bufferM, h.bufferN)

		// 2.4
		utils.AddInRing(h.bufferN, constants.V512)

		// 2,5
		utils.AddInRing(h.bufferSumm, h.bufferM)
	}

	return c, nil
}

func (h *Streebog) Sum(b []byte) []byte {
	// 3.1
	utils.BytesToUints(utils.PadBytes(h.bufferMessage), h.bufferM)

	copy(h.bufferOutH, h.bufferH)
	copy(h.bufferOutN, h.bufferN)
	copy(h.bufferOutSumm, h.bufferSumm)

	// 3.2
	round.G(h.bufferOutH, h.bufferM, h.bufferOutN)

	// 3.3
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(len(h.bufferMessage)*8))
	slices.Reverse(buf)
	utils.AddInRing(h.bufferOutN, []uint64{binary.BigEndian.Uint64(buf), 0, 0, 0, 0, 0, 0, 0})

	// 3.4
	utils.AddInRing(h.bufferOutSumm, h.bufferM)

	// 3.5
	round.G(h.bufferOutH, h.bufferOutN, constants.Zeroes)

	// 3.6
	round.G(h.bufferOutH, h.bufferOutSumm, constants.Zeroes)

	temp := utils.UintsToBytes(h.bufferOutH)
	if h.size == 32 {
		temp = temp[32:]
	}

	b = append(b, temp...)

	return b
}

func (h *Streebog) Reset() {
	h.bufferH = make([]uint64, 8)
	switch h.size {
	case 32:
		utils.BytesToUints(constants.IV256, h.bufferH)
	case 64:
		utils.BytesToUints(constants.IV512, h.bufferH)
	}
	h.bufferN = make([]uint64, 8)
	h.bufferM = make([]uint64, 8)
	h.bufferSumm = make([]uint64, 8)

	h.bufferOutH = make([]uint64, 8)
	h.bufferOutN = make([]uint64, 8)
	h.bufferOutSumm = make([]uint64, 8)
}
