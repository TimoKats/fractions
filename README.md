![fractions logo](https://github.com/TimoKats/fractions/blob/cafbfa70d4054a1df8eb4c10d2fc4d6511dd3f19/logo.png)  
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) 
[![Donate](https://img.shields.io/badge/Donate-PayPal-blue.svg)](https://www.paypal.com/donate/?hosted_button_id=V2YDND4NLLPRL)
![Status](https://img.shields.io/static/v1.svg?label=version&message=v1.0.1&color=blue)

---
### Getting started
Fractions is a Go library that adds fraction-like capabilities to Go. If you want to install Fractions simply run the command shown below in your terminal.  
``` go
go get github.com/TimoKats/fractions
```
Next, after installing Fractions you can import (and test) the installation with the code shown below. If this outputs `1 2/3` then the installation has been completed correctly.  
``` go
package main

import (
	"github.com/TimoKats/fractions"
)

func main() {
	y1 := fractions.FloatToFrac(1.67)
	fractions.PrintFrac(y1)
}
```
### Functionalities
This library saves fractions as objects/structs (called *frac*) and has functions for: printing, setting, getting, formatting and computing. Hereby a short overview of them all.  
#### print functions
 - PrintFrac(frac *frac)
 - PrintNumerator(frac *frac)
 - PrintDenominator(frac *frac)
 - PrintInteger(frac \*frac)
#### set functions
 - MakeFrac(integer int, numerator int, denominator int)
 - SetNumerator(frac \*frac, value int)
 - SetDenominator(frac \*frac, value int)
 - SetInteger(frac \*frac, value int)
#### get functions
 - GetNumerator(frac \*frac)
 - GetDenominator(frac \*frac)
 - GetInteger(frac \*frac)
#### conversion functions
 - FloatToFrac(value float64)
 - FracToFloat(frac \*frac)
#### formatting functions
 - FormatFracOnly(frac \*frac)
 - FormatSimplify(frac \*frac)
#### operations
 - SumFrac(frac1 \*frac, frac2 \*frac)
 - SumInt(frac1 \*frac, frac2 \*frac)
 - SubtractFrac(frac1 \*frac, frac2 \*frac)
 - SubtractInt(frac1 \*frac, frac2 \*frac)
 - SubtractInt(frac1 \*frac, value int)
 - DivideFrac(frac1 \*frac, frac2 \*frac)
 - DivideInt(frac1 \*frac, value int)
 - PowerFrac(frac1 \*frac, value int)
