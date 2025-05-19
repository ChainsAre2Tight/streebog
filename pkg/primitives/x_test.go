package primitives_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/streebog/pkg/primitives"
)

func TestX(t *testing.T) {
	tt := []struct {
		a, b, c []uint64
	}{
		{
			a: []uint64{0b0, 0b1, 0b11, 0b101, 0b10101, 0b100010, 0b11, 0b1011110},
			b: []uint64{0b0, 0b1, 0b10, 0b110, 0b11010, 0b100101, 0b01, 0b1101010},
			c: []uint64{0b0, 0b0, 0b01, 0b011, 0b01111, 0b000111, 0b10, 0b0110100},
		},
	}
	for _, td := range tt {
		t.Run(
			fmt.Sprintf("%0.16x ^ %0.16x -> %0.16x", td.a, td.b, td.c),
			func(t *testing.T) {
				primitives.X(td.a, td.b)
				if !reflect.DeepEqual(td.c, td.a) {
					t.Fatalf("\nGot:  %0.16x, \nWant: %0.16x.", td.a, td.c)
				}
			},
		)
	}
}
