package primitives_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog/pkg/primitives"
	"github.com/ChainsAre2Tight/streebog/pkg/testdata"
)

func TestL(t *testing.T) {
	tt := []struct {
		in  []uint64
		out []uint64
	}{
		{
			in:  testdata.PSXK1m,
			out: testdata.LPSXK1m,
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("%0.16x -> %0.16x", td.in, td.out),
			func(t *testing.T) {
				res := make([]uint64, 8)
				copy(res, td.in)
				primitives.L(res)
				if !reflect.DeepEqual(res, td.out) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x.", res, td.out)
				}
			},
		)
	}
}

func BenchmarkL(b *testing.B) {
	in := make([]uint64, 8)
	copy(in, testdata.PSXK1m)

	for b.Loop() {
		primitives.L(in)
	}
}
