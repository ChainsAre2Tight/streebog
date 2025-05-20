package round_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog/pkg/round"
	"github.com/ChainsAre2Tight/streebog/pkg/testdata"
)

func TestRoundFunction(t *testing.T) {
	tt := []struct {
		h, m, N, res []uint64
	}{
		{
			h:   testdata.Zero512,
			m:   testdata.M,
			N:   testdata.Zero512,
			res: testdata.Example1_G_result,
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("%0.16x / %0.16x / %0.16x -> %0.16x", td.h, td.m, td.N, td.res),
			func(t *testing.T) {
				res := make([]uint64, 8)
				copy(res, td.h)
				round.G(res, td.m, td.N)
				if !reflect.DeepEqual(res, td.res) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x.", res, td.res)
				}
			},
		)
	}
}

func BenchmarkRoundFunction(b *testing.B) {
	tt := []struct {
		h, m, N, res []uint64
	}{
		{
			h: testdata.Zero512,
			m: testdata.M,
			N: testdata.Zero512,
		},
	}
	for _, td := range tt {
		b.Run(
			fmt.Sprintf("%0.16x / %0.16x / %0.16x", td.h, td.m, td.N),
			func(b *testing.B) {
				in := make([]uint64, 8)
				for b.Loop() {
					round.G(in, td.m, td.N)
				}
			},
		)
	}
}
