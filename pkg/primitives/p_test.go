package primitives_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog/pkg/primitives"
	"github.com/ChainsAre2Tight/streebog/pkg/testdata"
)

func TestP(t *testing.T) {
	tt := []struct {
		in, out []uint64
	}{
		{
			in:  testdata.SXK1m,
			out: testdata.PSXK1m,
		}, {
			in:  testdata.SK1xC1,
			out: testdata.PSK1xC1,
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("%0.16x -> %0.16x", td.in, td.out),
			func(t *testing.T) {
				res := make([]uint64, 8)
				copy(res, td.in)
				primitives.P(res)
				if !reflect.DeepEqual(res, td.out) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x.", res, td.out)
				}
			},
		)
	}
}
