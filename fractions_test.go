// fractions_test.go
// written by Timo Kats
// last updated: 20/03/2022

package fractions_test

import (
	"testing"

	"github.com/TimoKats/fractions"
)

func TestFractions(t *testing.T) {
	// the answers I want to test
	Ey3 := 3.0
	Ey4 := 2.5

	// the operations I aim to test
	y1 := fractions.FloatToFrac(1.5)
	y2 := fractions.MakeFrac(0, 1, 2)

	y3 := fractions.DivideFrac(y1, y2)
	y4 := fractions.SumInt(y2, 2)

	// test the answers
	if Ey3 != fractions.FracToFloat(y3) || Ey4 != fractions.FracToFloat(y4) {
		t.Error("Calculated value not same as correct value.\n")
	}

}
