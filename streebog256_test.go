package streebog_test

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/ChainsAre2Tight/streebog"
	"github.com/ChainsAre2Tight/streebog/pkg/testdata"
)

func TestStreebog256(t *testing.T) {
	tt := []struct {
		in, out []byte
	}{
		{
			in:  testdata.M1,
			out: testdata.Case1_HASH_256,
		}, {
			in:  testdata.M2,
			out: testdata.Case2_HASH_256,
		},
	}
	for test, td := range tt {
		t.Run(
			fmt.Sprintf("test #%0.2d | %0.16x -> %0.16x", test, td.in[0:16], td.out[0:16]),
			func(t *testing.T) {
				slices.Reverse(td.in)
				res, err := streebog.Streebog256(td.in)
				slices.Reverse(td.out)
				if err != nil {
					t.Fatalf("error: %s", err)
				}
				if !reflect.DeepEqual(td.out, res) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x.", res, td.out)
				}
			},
		)
	}
}

func BenchmarkSteebog256(b *testing.B) {
	value := testdata.M2
	var hash []byte
	var err error
	b.ReportAllocs()

	for b.Loop() {
		hash, err = streebog.Streebog512(value)
		if err != nil {
			b.Fatalf("Error: %s", err)
		}
	}

	if len(hash) == 0 {
		b.Fatalf("hash is too short")
	}
}
