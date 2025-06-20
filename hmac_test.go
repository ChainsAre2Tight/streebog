package streebog_test

import (
	"crypto/hmac"
	"fmt"
	"hash"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog"
)

func TestHMAC256(t *testing.T) {
	tt := []struct {
		message, key, hmac256, hmac512 []byte
	}{
		{
			key: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			},
			message: []byte{
				0x01, 0x26, 0xbd, 0xb8, 0x78, 0x00, 0xaf, 0x21, 0x43, 0x41, 0x45, 0x65, 0x63, 0x78, 0x01, 0x00,
			},
			hmac256: []byte{
				0xa1, 0xaa, 0x5f, 0x7d, 0xe4, 0x02, 0xd7, 0xb3, 0xd3, 0x23, 0xf2, 0x99, 0x1c, 0x8d, 0x45, 0x34,
				0x01, 0x31, 0x37, 0x01, 0x0a, 0x83, 0x75, 0x4f, 0xd0, 0xaf, 0x6d, 0x7c, 0xd4, 0x92, 0x2e, 0xd9,
			},
			hmac512: []byte{
				0xa5, 0x9b, 0xab, 0x22, 0xec, 0xae, 0x19, 0xc6, 0x5f, 0xbd, 0xe6, 0xe5, 0xf4, 0xe9, 0xf5, 0xd8,
				0x54, 0x9d, 0x31, 0xf0, 0x37, 0xf9, 0xdf, 0x9b, 0x90, 0x55, 0x00, 0xe1, 0x71, 0x92, 0x3a, 0x77,
				0x3d, 0x5f, 0x15, 0x30, 0xf2, 0xed, 0x7e, 0x96, 0x4c, 0xb2, 0xee, 0xdc, 0x29, 0xe9, 0xad, 0x2f,
				0x3a, 0xfe, 0x93, 0xb2, 0x81, 0x4f, 0x79, 0xf5, 0x00, 0x0f, 0xfc, 0x03, 0x66, 0xc2, 0x51, 0xe6,
			},
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("256 %x | %x -> %x", td.message, td.key, td.hmac256),
			func(t *testing.T) {
				h := hmac.New(func() hash.Hash { return streebog.New(32) }, td.key)
				h.Write(td.message)
				res := h.Sum(nil)
				if !reflect.DeepEqual(res, td.hmac256) {
					t.Fatalf("\nGot:  %x, \nWant: %x.", res, td.hmac256)
				}
			},
		)
		t.Run(
			fmt.Sprintf("512 %x | %x -> %x", td.message, td.key, td.hmac512),
			func(t *testing.T) {
				h := hmac.New(func() hash.Hash { return streebog.New(64) }, td.key)
				h.Write(td.message)
				res := h.Sum(nil)
				if !reflect.DeepEqual(res, td.hmac512) {
					t.Fatalf("\nGot:  %x, \nWant: %x.", res, td.hmac512)
				}
			},
		)
	}
}
