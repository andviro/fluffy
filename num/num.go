// package num serves as a thin abstraction layer for numeric types. If
// nessessary we could replace it with alternative implementation using Go
// generics
package num

import (
	"strings"
)

type Num[T Num[T]] interface {
	Float() float64
	Cmp(a T) int
	NewF(float64) T
	NewS(string) T
	NewI(int64, uint) T
	Sqrt() T
	ZERO() T
	Abs() T
	NaN() T
	Exp(T) T
	Neg() T
	Sub(T) T
	Div(T) T
	Mul(T) T
	Add(T) T
	IsZero() bool
	IsNaN() bool
	LessThan(T) bool
	GreaterThan(T) bool
	LessThanOrEqual(T) bool
	Equal(T) bool
	GreaterThanOrEqual(T) bool
	Sign() int
	Int() int64
	String() string
}

func StringN[T Num[T]](f T, decimals int) string {
	if f.IsNaN() {
		return "NaN"
	}
	s := f.String()
	if decimals < 0 {
		return s
	}
	ss := strings.Split(f.String(), ".")
	if len(ss) == 1 {
		s += "."
		ss = append(ss, ".")
	}
	point := len(ss[0])
	if decimals == 0 {
		return s[:point]
	} else {
		for i := point + len(ss[1]); i < point+decimals+1; i++ {
			s += "0"
		}
		return s[:point+decimals+1]
	}
}

func ZERO[T Num[T]]() (res T) {
	return res
}

func NewI[T Num[T]](a int64, b uint) (res T) {
	return res.NewI(a, b)
}

func NewS[T Num[T]](s string) (res T) {
	return res.NewS(s)
}

func NewF[T Num[T]](a float64) (res T) {
	return res.NewF(a)
}

func One[T Num[T]]() (res T) {
	return res.NewI(1, 0)
}

func Two[T Num[T]]() (res T) {
	return res.NewI(2, 0)
}

func Sqr[T Num[T]](x T) (res T) {
	return x.Mul(x)
}

func Neg[T Num[T]](x T) (res T) {
	return x.Neg()
}

func NaN[T Num[T]]() (res T) {
	return res.NaN()
}

func Hundred[T Num[T]]() (res T) {
	return res.NewI(100, 0)
}

func Max[T Num[T]](first T, rest ...T) T {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) > 0 {
			ans = item
		}
	}
	return ans
}

func Min[T Num[T]](first T, rest ...T) T {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) < 0 {
			ans = item
		}
	}
	return ans
}
