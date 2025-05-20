package streebog_test

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog"
)

func d(in string) []byte {
	res, err := hex.DecodeString(in)
	if err != nil {
		panic(fmt.Errorf("streebog_test.d: %s", err))
	}
	return res
}

func TestStreebog512(t *testing.T) {
	tt := []struct {
		in, out []byte
	}{
		{
			in:  d("fbe2e5f0eee3c820fbeafaebef20fffbf0e1e0f0f520e0ed20e8ece0ebe5f0f2f120fff0eeec20f120faf2fee5e2202ce8f6f3ede220e8e6eee1e8f0f2d1202ce8f0f2e5e220e5d1"),
			out: d("28fbc9bada033b1460642bdcddb90c3fb3e56c497ccd0f62b8a2ad4935e85f037613966de4ee00531ae60f3b5a47f8dae06915d5f2f194996fcabf2622e6881e"),
		},
	}
	for test, td := range tt {
		t.Run(
			fmt.Sprintf("test #%0.2d | %0.16x -> %0.16x", test, td.in[0:16], td.out[0:16]),
			func(t *testing.T) {
				res, err := streebog.Streebog512(td.in)
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
