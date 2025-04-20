package fixed

import (
	"math"

	"github.com/andviro/fluffy/v2/num"
	"github.com/robaho/fixed"
)

type Fixed fixed.Fixed

var ZERO Fixed

// Cmp implements num.Num.
func (f Fixed) Cmp(a Fixed) int {
	return fixed.Fixed(f).Cmp(fixed.Fixed(a))
}

// Float implements num.Num.
func (f Fixed) Float() float64 {
	return (fixed.Fixed)(f).Float()
}

// NaN implements num.Num.
func (f Fixed) NaN() Fixed {
	return Fixed(fixed.NaN)
}

// NewF implements num.Num.
func (f Fixed) NewF(val float64) Fixed {
	return Fixed(fixed.NewF(val))
}

// NewS implements num.Num.
func (f Fixed) NewS(val string) Fixed {
	return Fixed(fixed.NewS(val))
}

// NewI implements num.Num.
func (f Fixed) NewI(val int64, n uint) Fixed {
	return Fixed(fixed.NewI(val, n))
}

// Sqrt implements num.Num.
func (f Fixed) Sqrt() Fixed {
	return Fixed(fixed.NewF(math.Sqrt(f.Float())))
}

// ZERO implements num.Num.
func (f Fixed) ZERO() Fixed {
	return Fixed(fixed.ZERO)
}

// IsZero implements num.Num.
func (f Fixed) IsZero() bool {
	return fixed.Fixed(f).IsZero()
}

// IsNan implements num.Num.
func (f Fixed) IsNaN() bool {
	return fixed.Fixed(f).IsNaN()
}

// Abs implements num.Num.
func (f Fixed) Abs() Fixed {
	return Fixed(fixed.Fixed(f).Abs())
}

func (f Fixed) Exp(n Fixed) Fixed {
	return f.NewF(math.Exp(n.Float()))
}

func (f Fixed) Sub(n Fixed) Fixed {
	return Fixed(fixed.Fixed(f).Sub(fixed.Fixed(n)))
}

func (f Fixed) Add(n Fixed) Fixed {
	return Fixed(fixed.Fixed(f).Add(fixed.Fixed(n)))
}

func (f Fixed) Div(n Fixed) Fixed {
	return Fixed(fixed.Fixed(f).Div(fixed.Fixed(n)))
}

func (f Fixed) Mul(n Fixed) Fixed {
	return Fixed(fixed.Fixed(f).Mul(fixed.Fixed(n)))
}

func (f Fixed) LessThan(n Fixed) bool {
	return fixed.Fixed(f).LessThan(fixed.Fixed(n))
}

func (f Fixed) GreaterThan(n Fixed) bool {
	return fixed.Fixed(f).GreaterThan(fixed.Fixed(n))
}

func (f Fixed) LessThanOrEqual(n Fixed) bool {
	return fixed.Fixed(f).LessThanOrEqual(fixed.Fixed(n))
}

func (f Fixed) GreaterThanOrEqual(n Fixed) bool {
	return fixed.Fixed(f).GreaterThanOrEqual(fixed.Fixed(n))
}

func (f Fixed) Equal(n Fixed) bool {
	return fixed.Fixed(f).Equal(fixed.Fixed(n))
}

func (f Fixed) Neg() Fixed {
	return f.ZERO().Sub(f)
}

func (f Fixed) Sign() int {
	return fixed.Fixed(f).Sign()
}

func (f Fixed) Int() int64 {
	return fixed.Fixed(f).Int()
}

func (f Fixed) String() string {
	return fixed.Fixed(f).String()
}

func (f *Fixed) UnmarshalJSON(bytes []byte) error {
	return (*fixed.Fixed)(f).UnmarshalJSON(bytes)
}

func (f Fixed) MarshalJSON() ([]byte, error) {
	return fixed.Fixed(f).MarshalJSON()
}

func (f Fixed) MarshalBinary() ([]byte, error) {
	return fixed.Fixed(f).MarshalBinary()
}

func (f *Fixed) UnmarshalBinary(bytes []byte) error {
	return (*fixed.Fixed)(f).UnmarshalBinary(bytes)
}

var _ num.Num[Fixed] = Fixed{}
