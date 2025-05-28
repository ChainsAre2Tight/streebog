package streebog_test

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
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
		}, {
			in:   []byte{'S', 't', 'r', 'e', 'e', 'b', 'o', 'g', 'T', 'e', 's', 't', 'I', 'n', 'p', 'u', 't', 'S', 't', 'r', 'i', 'n', 'g'},
			h256: []byte{0x77, 0x6d, 0xb6, 0x11, 0xfb, 0xea, 0xc2, 0xfe, 0x73, 0xb3, 0xf4, 0xb4, 0x83, 0x30, 0x5e, 0x46, 0x7a, 0x30, 0x4e, 0x51, 0x6f, 0xb0, 0xca, 0xe0, 0xcb, 0x20, 0x7f, 0xcd, 0xa4, 0xe7, 0x7c, 0xa1},
			h512: []byte{0x99, 0x63, 0x91, 0x41, 0x21, 0x93, 0x3c, 0xd0, 0x01, 0x14, 0x98, 0x36, 0x38, 0xd5, 0x94, 0x86, 0x04, 0x8a, 0xdf, 0x59, 0x6d, 0x74, 0xb4, 0xc6, 0x0a, 0xa5, 0x1b, 0xab, 0x64, 0x49, 0x22, 0xef, 0x30, 0xb7, 0x56, 0xda, 0x55, 0x84, 0x9e, 0x06, 0x78, 0x76, 0x90, 0x96, 0x39, 0x34, 0xfa, 0x77, 0xf3, 0x1c, 0xed, 0xcc, 0x13, 0x57, 0x5c, 0x44, 0x43, 0xcc, 0xc0, 0x71, 0xe6, 0x33, 0xdf, 0x2b},
		}, {
			in:   []byte("The quick brown fox jumps over the lazy dog"),
			h256: d("3e7dea7f2384b6c5a3d0e24aaa29c05e89ddd762145030ec22c71a6db8b2c1f4"),
			h512: d("d2b793a0bb6cb5904828b5b6dcfb443bb8f33efc06ad09368878ae4cdc8245b97e60802469bed1e7c21a64ff0b179a6a1e0bb74d92965450a0adab69162c00fe"),
		}, {
			in:   []byte("The quick brown fox jumps over the lazy dog."),
			h256: d("36816a824dcbe7d6171aa58500741f2ea2757ae2e1784ab72c5c3c6c198d71da"),
			h512: d("fe0c42f267d921f940faa72bd9fcf84f9f1bd7e9d055e9816e4c2ace1ec83be82d2957cd59b86e123d8f5adee80b3ca08a017599a9fc1a14d940cf87c77df070"),
		}, {
			in:   d("a2122f4c71550761c6b6c81ae04b0e0e082d36bb46d98ce7d257cbb2a0f51fd266d3a3b7db069782df2446d7c58960347d6ed4fbc82acf6b6cb2b253e833ecb3e712ba61642d89e68cb4983745adc13c19c3dcb88f2d68f895c44b189b7783411fa1684d572e0896914a39d9046075a9a96f52041fcf071360c6c61e6097f18ae01a0abe5bc59246dcc616383e8876c8b6662e8890dffaee007b20f9a99e3905e73f218070e58fbe1cc2132966d423bf7acb2c764981078d38d360d5f1e8b0c161bf190546433b4457edb83281d6fa501b9776554f0e358e9fcf92a4cdb31315a43313c561952e8cc9b52832381847876661b78c6ec295fcef65a0817b6ec24d"),
			h256: d("65920ee24757a07d9b7b9c038e01f81faa4a317d950a10c47e286efea05df794"),
			h512: d("2d8a9f09206e95dda2fa2050afaff0f8de6d161cb0f76d53589c7ce3f55e1539448d7a9ea27c174ee6459ad07317136adad5cd9e9fef7dcb7260be0cd41f1efe"),
		},
	}
	for test, td := range tt {
		t.Run(
			fmt.Sprintf("test #%0.2d | %0.16x -> %0.16x [h256]", test, td.in[0:16], td.h256[0:16]),
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
			},
		)
		t.Run(
			fmt.Sprintf("test #%0.2d | %0.16x -> %0.16x [h512]", test, td.in[0:16], td.h512[0:16]),
			func(t *testing.T) {
				h := streebog.New(64)
				_, err := h.Write(td.in)
				res := h.Sum(nil)

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

func TestEmpty(t *testing.T) {
	t.Run("256", func(t *testing.T) {
		h := streebog.New(32)
		h.Write([]byte{})
		if !bytes.Equal(h.Sum(nil),
			d("3f539a213e97c802cc229d474c6aa32a825a360b2a933a949fd925208d9ce1bb"),
		) {
			t.FailNow()
		}
	})
	t.Run("512", func(t *testing.T) {
		h := streebog.New(64)
		h.Write([]byte{})
		if !bytes.Equal(h.Sum(nil),
			d("8e945da209aa869f0455928529bcae4679e9873ab707b55315f56ceb98bef0a7362f715528356ee83cda5f2aac4c6ad2ba3a715c1bcd81cb8e9f90bf4c1c1a8a"),
		) {
			t.FailNow()
		}
	})
}

func TestBlocksized(t *testing.T) {
	m := make([]byte, 64)
	for i := range 64 {
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

func BenchmarkStreebogWrite(b *testing.B) {
	h := streebog.New(64)
	src := make([]byte, 65)
	rand.Read(src)

	for b.Loop() {
		h.Write(src)
	}
}

func BenchmarkStreebogSum(b *testing.B) {
	h := streebog.New(64)
	src := make([]byte, 65)
	rand.Read(src)
	h.Write(src)

	r := make([]byte, 0, 64)

	for b.Loop() {
		h.Sum(r)
		r = r[:0]
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

	t.Run("Sum with non-nil argument", func(t *testing.T) {
		h := streebog.New(64)
		blocks := make([]byte, 96)
		for i := range blocks {
			blocks[i] = byte(i)
		}
		h.Write(blocks[:96])
		expected := h.Sum(nil)
		prefix := []byte{0xAA, 0xBB}
		res := h.Sum(prefix)
		if !bytes.Equal(res[:2], prefix) {
			t.Fatalf("prefix not preserved in Sum: got %x, want prefix %x", res[:2], prefix)
		}
		if !bytes.Equal(res[2:], expected) {
			t.Fatalf("hash not appended correctly in Sum with non-nil arg\nGot:  %0.16x, \nWant: %0.16x.", res[2:], expected)
		}
	})
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

	t.Run("96x1b", func(t *testing.T) {
		h := streebog.New(64)
		for i := range 96 {
			h.Write(blocks[i : i+1])
		}
		res := h.Sum(nil)
		if !bytes.Equal(res, expected) {
			t.FailNow()
		}
	})
}

func TestStreebogSize(t *testing.T) {
	t.Run("size 32", func(t *testing.T) {
		h := streebog.New(32)
		if h.Size() != 32 {
			t.Fatalf("expected Size() == 32, got %d", h.Size())
		}
	})
	t.Run("size 64", func(t *testing.T) {
		h := streebog.New(64)
		if h.Size() != 64 {
			t.Fatalf("expected Size() == 64, got %d", h.Size())
		}
	})
}

func TestStreebogNew(t *testing.T) {
	t.Run("valid size 32", func(t *testing.T) {
		h := streebog.New(32)
		if h == nil {
			t.Fatal("streebog.New(32) returned nil")
		}
		if h.Size() != 32 {
			t.Fatalf("expected Size() == 32, got %d", h.Size())
		}
	})
	t.Run("valid size 64", func(t *testing.T) {
		h := streebog.New(64)
		if h == nil {
			t.Fatal("streebog.New(64) returned nil")
		}
		if h.Size() != 64 {
			t.Fatalf("expected Size() == 64, got %d", h.Size())
		}
	})
	t.Run("invalid size panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic for invalid size, but did not panic")
			}
		}()
		_ = streebog.New(48)
	})
}

func TestBehaviourChunked48(t *testing.T) {
	msg := make([]byte, 48)
	for i := range msg {
		msg[i] = byte(i)
	}

	h := streebog.New(64)
	h.Write(msg)
	expected := h.Sum(nil)

	t.Run("2x24", func(t *testing.T) {
		h := streebog.New(64)
		h.Write(msg[:24])
		h.Write(msg[24:])
		res := h.Sum(nil)
		if !bytes.Equal(res, expected) {
			t.Fatalf("hash mismatch for 2x24: got %x, want %x", res, expected)
		}
	})

	t.Run("4x12", func(t *testing.T) {
		h := streebog.New(64)
		for i := 0; i < 48; i += 12 {
			h.Write(msg[i : i+12])
		}
		res := h.Sum(nil)
		if !bytes.Equal(res, expected) {
			t.Fatalf("hash mismatch for 4x12: got %x, want %x", res, expected)
		}
	})

	t.Run("3x16", func(t *testing.T) {
		h := streebog.New(64)
		for i := 0; i < 48; i += 16 {
			h.Write(msg[i : i+16])
		}
		res := h.Sum(nil)
		if !bytes.Equal(res, expected) {
			t.Fatalf("hash mismatch for 3x16: got %x, want %x", res, expected)
		}
	})
}
func d(in string) []byte {
	res, err := hex.DecodeString(in)
	if err != nil {
		panic(fmt.Errorf("streebog_test.d: %s", err))
	}
	return res
}
