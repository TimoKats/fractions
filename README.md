![fractions logo](https://github.com/TimoKats/fractions/blob/cafbfa70d4054a1df8eb4c10d2fc4d6511dd3f19/logo.png)  
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) 
[![Donate](https://img.shields.io/badge/Donate-PayPal-blue.svg)](https://www.paypal.com/donate/?hosted_button_id=V2YDND4NLLPRL)

---
### Getting started
Fractions is a Go library that adds fraction-like capabilities to Go. If you want to install Fractions simply run the command shown below in your terminal.  
```
go get github.com/TimoKats/fractions
```
Next, after installing Fractions you can import (and test) the installation with the code shown below. If this outputs `1 2/3` then the installation has been completed correctly.  
```
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

### Contact and donate
