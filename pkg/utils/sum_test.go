package utils_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog/pkg/constants"
	"github.com/ChainsAre2Tight/streebog/pkg/testdata"
	"github.com/ChainsAre2Tight/streebog/pkg/utils"
)

func TestSumInRing(t *testing.T) {
	tt := []struct {
		dst, add, res []uint64
	}{
		{
			dst: constants.Zeroes,
			add: testdata.M,
			res: testdata.M,
		}, { // TODO: add case 2 test data to testdata package and write a test case for summing in step 3.4
			dst: testdata.Case2_EPSILON_1,
			add: testdata.M2_chunk_2,
			res: testdata.Case2_EPSILON_2,
		},
	}
	for test, td := range tt {
		t.Run(
			fmt.Sprintf("test #%d | %0.16x + %0.16x -> %0.16x", test, td.dst[0], td.add[0], td.res[0]),
			func(t *testing.T) {
				res := make([]uint64, 8)
				copy(res, td.dst)
				utils.AddInRing(res, td.add)
				if !reflect.DeepEqual(res, td.res) {
					t.Fatalf("\nGot:  %0.16x,\nWant: %0.16x.", res, td.res)
				}
			},
		)
	}
}
