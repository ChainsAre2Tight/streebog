package primitives_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog/pkg/primitives"
	"github.com/ChainsAre2Tight/streebog/pkg/testdata"
)

func TestS(t *testing.T) {
	tt := []struct {
		in  []uint64
		out []uint64
	}{
		{
			in:  testdata.XK1m,
			out: testdata.SXK1m,
		}, {
			in:  testdata.K1xC1,
			out: testdata.SK1xC1,
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("%0.16x -> %0.16x", td.in, td.out),
			func(t *testing.T) {
				res := make([]uint64, 8)
				copy(res, td.in)
				primitives.S(res)
				if !reflect.DeepEqual(res, td.out) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x", res, td.out)
				}
			},
		)
	}
}

func BenchmarkS(b *testing.B) {
	in := make([]uint64, 8)
	copy(in, testdata.XK1m)

	for b.Loop() {
		primitives.S(in)
	}
}
