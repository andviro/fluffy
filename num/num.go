// package num serves as a thin abstraction layer for numeric types. If
// nessessary we could replace it with alternative implementation using Go
// modules system
package num

import (
	"math"

	num "github.com/robaho/fixed"
)

type Num = num.Fixed

var (
	NewI = num.NewI
	NewF = num.NewF
	NewS = num.NewS
	ZERO = num.ZERO
	NaN  = num.NaN
)

func Sqrt(val Num) Num {
	return num.NewF(math.Sqrt(val.Float()))
}

func Max(first Num, rest ...Num) Num {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) > 0 {
			ans = item
		}
	}
	return ans
}

func Min(first Num, rest ...Num) Num {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) < 0 {
			ans = item
		}
	}
	return ans
}

func Exp(n Num) Num {
	return NewF(math.Exp(n.Float()))
}

func Neg(n Num) Num {
	return ZERO.Sub(n)
}
