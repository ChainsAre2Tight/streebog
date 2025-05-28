package streebog

import (
	"encoding/binary"
	"hash"
	"slices"

	"github.com/ChainsAre2Tight/streebog/pkg/constants"
	"github.com/ChainsAre2Tight/streebog/pkg/round"
	"github.com/ChainsAre2Tight/streebog/pkg/utils"
)

var _ hash.Hash = (*streebog)(nil)

type streebog struct {
	size int

	// used to store message chunks as they are passed into Write()
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

func (h *streebog) BlockSize() int {
	return 64
}

func (h *streebog) Size() int {
	return h.size
}
func New(size int) hash.Hash {
	if size != 32 && size != 64 {
		panic("invalid hash size")
	}
	res := &streebog{
		size: size,
	}
	res.Reset()
	return res
}

func (h *streebog) Write(p []byte) (n int, err error) {

	// got remembers the length of p to return it later
	got := len(p)

	fillBuffer := func() {
		remaining := 64 - len(h.bufferMessage)

		// this should never reallocate, cap(h.bufferMessage) must remain at 64
		if remaining < len(p) {
			h.bufferMessage = append(h.bufferMessage, p[:remaining]...)
			p = p[remaining:]
		} else {
			h.bufferMessage = append(h.bufferMessage, p...)
			p = p[len(p):]
		}
	}

	fillBuffer()

	// 2.1
	// if length == 64 we have enough data in p to fill the buffer
	for len(h.bufferMessage) == 64 {

		// 2.2
		utils.BytesToUints(h.bufferMessage, h.bufferM)
		h.bufferMessage = h.bufferMessage[:0]
		fillBuffer()

		// 2.3
		round.G(h.bufferH, h.bufferM, h.bufferN)

		// 2.4
		utils.AddInRing(h.bufferN, constants.V512)

		// 2.5
		utils.AddInRing(h.bufferSumm, h.bufferM)
	}

	return got, nil
}

func (h *streebog) Sum(b []byte) []byte {
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

func (h *streebog) Reset() {
	h.bufferH = make([]uint64, 8)
	switch h.size {
	case 32:
		utils.BytesToUints(constants.IV256, h.bufferH)
	case 64:
		utils.BytesToUints(constants.IV512, h.bufferH)
	}
	h.bufferMessage = make([]byte, 0, 64)

	h.bufferN = make([]uint64, 8)
	h.bufferM = make([]uint64, 8)
	h.bufferSumm = make([]uint64, 8)

	h.bufferOutH = make([]uint64, 8)
	h.bufferOutN = make([]uint64, 8)
	h.bufferOutSumm = make([]uint64, 8)
}
