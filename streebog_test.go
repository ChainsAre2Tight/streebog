package streebog_test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/streebog"
	"github.com/ChainsAre2Tight/streebog/pkg/testdata"
)

func TestStreebog(t *testing.T) {
	tt := []struct {
		in, h256, h512 []byte
	}{
		{
			in:   testdata.M1,
			h256: testdata.Case1_HASH_256,
			h512: testdata.Case1_HASH_512,
		}, {
			in:   testdata.M2,
			h256: testdata.Case2_HASH_256,
			h512: testdata.Case2_HASH_512,
		},
	}
	for test, td := range tt {
		t.Run(
			fmt.Sprintf("test #%0.2d | %0.16x -> %0.16x", test, td.in[0:16], td.h256[0:16]),
			func(t *testing.T) {
				h := streebog.New(32)
				_, err := h.Write(td.in)
				res := h.Sum(nil)

				if err != nil {
					t.Fatalf("error 256: %s", err)
				}
				if !bytes.Equal(td.h256, res) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x.", res, td.h256)
				}

				h = streebog.New(64)
				_, err = h.Write(td.in)
				res = h.Sum(nil)

				if err != nil {
					t.Fatalf("error 512: %s", err)
				}
				if !bytes.Equal(td.h512, res) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x.", res, td.h512)
				}
			},
		)
	}
}

func TestBlocksized(t *testing.T) {
	m := make([]byte, 64)
	for i := 0; i < 64; i++ {
		m[i] = byte(i)
	}
	h := streebog.New(64)
	h.Write(m)
	if !bytes.Equal(h.Sum(nil), []byte{
		0x2a, 0xe5, 0x81, 0xf1, 0x8a, 0xe8, 0x5e, 0x35,
		0x96, 0xc9, 0x36, 0xac, 0xbe, 0xf9, 0x10, 0xf2,
		0xed, 0x70, 0xdc, 0xf9, 0x1e, 0xd5, 0xd2, 0x4b,
		0x39, 0xa5, 0xaf, 0x65, 0x7b, 0xf8, 0x23, 0x2a,
		0x30, 0x3d, 0x68, 0x60, 0x56, 0xc8, 0xc0, 0x0b,
		0xf3, 0x0d, 0x42, 0xe1, 0x6c, 0xe2, 0x55, 0x42,
		0x6f, 0xa8, 0xa1, 0x55, 0xdc, 0xb3, 0xeb, 0x82,
		0x2d, 0x92, 0x58, 0x08, 0xf7, 0xc7, 0xe3, 0x45,
	}) {
		t.FailNow()
	}
}

func BenchmarkStreebog(b *testing.B) {
	h := streebog.New(64)
	src := make([]byte, 65)
	rand.Read(src)

	for b.Loop() {
		h.Write(src)
		h.Sum(nil)
	}
}

func TestBehaviour(t *testing.T) {
	h := streebog.New(64)

	// Sum does not change the state
	hsh1 := h.Sum(nil)
	if !bytes.Equal(h.Sum(nil), hsh1) {
		t.FailNow()
	}

	// No data equals to no state changing
	h.Write([]byte{})
	if !bytes.Equal(h.Sum(nil), hsh1) {
		t.FailNow()
	}

	// Just to be sure
	h.Write([]byte{})
	if !bytes.Equal(h.Sum(nil), hsh1) {
		t.FailNow()
	}
}

func TestBehaviourAdvanced(t *testing.T) {
	// fill block with distinct values
	// block contains 1.5 blocks of distinct, ordered bytes
	blocks := make([]byte, 96)
	for i := range blocks {
		blocks[i] = byte(i)
	}

	// write a message that is 1.5 blocks long
	// remember sum
	h := streebog.New(64)
	h.Write(blocks[:96])
	expected := h.Sum(nil)

	t.Run("1.5 blocks", func(t *testing.T) {
		// write a message that is 1 block long
		// and then 0.5 blocks long
		// sum must be equal to the first one
		h := streebog.New(64)
		h.Write(blocks[:64])
		h.Write(blocks[64:96])
		res := h.Sum(nil)

		if !bytes.Equal(res, expected) {
			t.FailNow()
		}

	})

	t.Run("0.75+0.75", func(t *testing.T) {
		// write a 0.75 block long message twice
		// sum must be equal to the first one
		h := streebog.New(64)
		h.Write(blocks[:48])
		h.Write(blocks[48:96])
		res := h.Sum(nil)

		if !bytes.Equal(res, expected) {
			t.FailNow()
		}
	})

	t.Run("0.5+0.5+0.5", func(t *testing.T) {
		// write a 0.5 block long message thrice
		// sum must be equal to the first one
		h := streebog.New(64)
		h.Write(blocks[:32])
		h.Write(blocks[32:64])
		h.Write(blocks[64:96])
		res := h.Sum(nil)

		if !bytes.Equal(res, expected) {
			t.FailNow()
		}
	})
}
