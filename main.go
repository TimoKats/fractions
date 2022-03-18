package main

import (
	"fmt"
)

// non-fraction functions //

func package_info() {
	fmt.Println("package-name:	fractions")
	fmt.Println("version:		0.0.1")
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

	print("\nERROR: no viable fractions found...\n")
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
	if frac.numerator > frac.denominator {
		frac.integer += frac.numerator / frac.denominator
		frac.numerator %= frac.denominator
	}
}

// operations

func MultiplyFrac(frac1 *frac, frac2 *frac) *frac {
	FormatFracOnly(frac1)
	FormatFracOnly(frac2)

	numerator := frac1.numerator * frac2.numerator
	denominator := frac1.denominator * frac2.denominator

	return &frac{0, numerator, denominator}
}

func MultiplyNum(frac1 *frac, value int) *frac {
	FormatFracOnly(frac1)

	numerator := frac1.numerator * value

	return &frac{0, numerator, frac1.denominator}
}

func main() {
	//package_info()

	y1 := FloatToFrac(2.5)
	//y2 := MakeFrac(0, 1, 2)

	y3 := MultiplyNum(y1, 2)
	FormatSimplify(y3)
	PrintFrac(y3)
}
