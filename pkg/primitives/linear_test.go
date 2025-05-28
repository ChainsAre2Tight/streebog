package primitives

import (
	"fmt"
	"testing"
)

func TestLinear(t *testing.T) {
	tt := []struct {
		in  uint64
		out uint64
	}{
		{
			// 	in:  0x46433ed624df433e,
			// 	out: 0xe60059d4d8e07580,
			// }, {
			in:  0x3e43df24d63e4346,
			out: 0x8075e0d8d45900e6,
		},
	}
	for num, td := range tt {
		t.Run(
			fmt.Sprintf("Test %02d | %0.16x -> %0.16x", num, td.in, td.out),
			func(t *testing.T) {
				res := linear(td.in)
				if res != td.out {
					t.Fatalf("\nGot:  %x, \nWant: %x", res, td.out)
				}
			},
		)
	}
}
