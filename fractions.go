// fractions.go
// written by Timo Kats
// last updated: 20/03/2022

package fractions

import (
	"fmt"
	"math"
)

// non-fraction functions //

func package_info() {
	fmt.Println("package-name:	fractions")
	fmt.Println("version:		v1.0.0")
	fmt.Println("made by:		Timo Kats")
	fmt.Println("last update:	17/03/2022")
}

// fraction functions //

type frac struct {
	integer, numerator, denominator int
}

// print functions //

func PrintFrac(frac *frac) {
	if frac.integer == 0 {
		if frac.numerator != 0 {
			print(frac.numerator, "/", frac.denominator)
		} else {
			print("ERROR: no values in fraction")
		}
	} else {
		if frac.numerator != 0 {
			print(frac.integer, " ", frac.numerator, "/", frac.denominator)
		} else {
			print(frac.integer)
		}
	}
	print("\n")
}

func PrintNumerator(frac *frac) {
	print(frac.numerator)
}

func PrintDenominator(frac *frac) {
	print(frac.denominator)
}

func PrintInteger(frac *frac) {
	print(frac.integer)
}

// set functions //

func SetNumerator(frac *frac, value int) {
	frac.numerator = value
}

func SetDenominator(frac *frac, value int) {
	frac.denominator = value
}

func SetInteger(frac *frac, value int) {
	frac.integer = value
}

func MakeFrac(integer int, numerator int, denominator int) *frac {
	NewFrac := &frac{integer, numerator, denominator}
	return NewFrac
}

// get functions //

func GetNumerator(frac *frac) int {
	return frac.numerator
}

func GetDenominator(frac *frac) int {
	return frac.denominator
}

func GetInteger(frac *frac) int {
	return frac.integer
}

// conversion functions //

func FloatToFrac(value float64) *frac {
	integer := int(value)
	decimal := value - float64(integer)

	for d := 1; d < 100; d++ {
		for n := 1; n < d; n++ {
			if (float64(n)/float64(d)) <= (decimal+0.01) && (float64(n)/float64(d)) >= (decimal-0.01) {
				return &frac{integer, n, d}
			}
		}
	}
	if value >= 0 {
		print("\nERROR: no viable fractions found\n")
	} else {
		print("\nERROR: no non-positive values allowed\n")
	}
	return &frac{0, 0, 0}
}

func FracToFloat(frac *frac) float64 {
	return float64(frac.numerator)/float64(frac.denominator) + float64(frac.integer)
}

// formatting functions //

func FormatFracOnly(frac *frac) {
	frac.numerator += frac.denominator * frac.integer
	frac.integer = 0
}

func FormatSimplify(frac *frac) {
	if frac.numerator >= frac.denominator {
		frac.integer += frac.numerator / frac.denominator
		frac.numerator %= frac.denominator
	} else if frac.numerator < 0 && math.Abs(float64(frac.numerator)) >= float64(frac.denominator) {
		frac.integer += frac.numerator / frac.denominator
		frac.numerator %= frac.denominator
	} else {
		print("ERROR: frac already simplified\n")
	}
}

// boolean checks

func EqualDenominators(frac1 *frac, frac2 *frac) bool {
	if frac1.denominator == frac2.denominator {
		return true
	} else {
		return false
	}
}

// operations

func EqualizeDenominators(frac1 *frac, frac2 *frac) {
	if !EqualDenominators(frac1, frac2) {
		FormatFracOnly(frac1)
		FormatFracOnly(frac2)

		CommonDenominator := frac1.denominator * frac2.denominator
		frac1.numerator *= frac2.denominator
		frac2.numerator *= frac1.denominator
		frac1.denominator, frac2.denominator = CommonDenominator, CommonDenominator
	}
}

func MultiplyFrac(frac1 *frac, frac2 *frac) *frac {
	FormatFracOnly(frac1)
	FormatFracOnly(frac2)

	numerator := frac1.numerator * frac2.numerator
	denominator := frac1.denominator * frac2.denominator
	return &frac{0, numerator, denominator}
}

func MultiplyInt(frac1 *frac, value int) *frac {
	FormatFracOnly(frac1)

	numerator := frac1.numerator * value
	return &frac{0, numerator, frac1.denominator}
}

func SumFrac(frac1 *frac, frac2 *frac) *frac {
	FormatFracOnly(frac1)
	FormatFracOnly(frac2)
	EqualizeDenominators(frac1, frac2)

	numerator := frac1.numerator + frac2.numerator
	return &frac{0, numerator, frac1.denominator}
}

func SumInt(frac1 *frac, value int) *frac {
	integer := frac1.integer + value
	return &frac{integer, frac1.numerator, frac1.denominator}
}

func SubtractFrac(frac1 *frac, frac2 *frac) *frac {
	FormatFracOnly(frac1)
	FormatFracOnly(frac2)
	EqualizeDenominators(frac1, frac2)

	numerator := frac1.numerator - frac2.numerator
	return &frac{0, numerator, frac1.denominator}
}

func SubtractInt(frac1 *frac, value int) *frac {
	FormatFracOnly(frac1)
	value *= frac1.denominator

	numerator := frac1.numerator - value
	return &frac{0, numerator, frac1.denominator}
}

func DivideFrac(frac1 *frac, frac2 *frac) *frac {
	FormatFracOnly(frac1)
	FormatFracOnly(frac2)

	numerator := frac1.numerator * frac2.denominator
	denominator := frac1.denominator * frac2.numerator
	return &frac{0, numerator, denominator}
}

func DivideInt(frac1 *frac, value int) *frac {
	FormatFracOnly(frac1)
	denominator := frac1.denominator * value
	return &frac{0, frac1.numerator, denominator}
}

func PowerFrac(frac1 *frac, value int) *frac {
	FloatValue := math.Pow(FracToFloat(frac1), float64(value))
	return FloatToFrac(FloatValue)
}
