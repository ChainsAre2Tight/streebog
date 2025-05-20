package streebog_test

import (
	"fmt"
	"reflect"
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
				res, err := streebog.Streebog256(td.in)
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
