package float

import (
	"cmp"
	"fmt"
	"math"
	"strconv"
)

type Float float64

// Cmp implements num.Num.
func (f Float) Cmp(a Float) int {
	return cmp.Compare(f, a)
}

// Float implements num.Num.
func (f Float) Float() float64 {
	return float64(f)
}

// Int implements num.Num.
func (f Float) Int() int64 {
	return int64(f)
}

// NaN implements num.Num.
func (f Float) NaN() Float {
	return Float(math.NaN())
}

// NewF implements num.Num.
func (f Float) NewF(val float64) Float {
	return Float(val)
}

// NewS implements num.Num.
func (f Float) NewS(val string) Float {
	res, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return f.NaN()
	}
	return Float(res)
}

// NewI implements num.Num.
func (f Float) NewI(val int64, n uint) Float {
	return f.NewF(float64(val) / math.Pow10(int(n)))
}

// Sqrt implements num.Num.
func (f Float) Sqrt() Float {
	return Float(math.Sqrt(float64(f)))
}

// ZERO implements num.Num.
func (f Float) ZERO() Float {
	return Float(.0)
}

// IsZero implements num.Num.
func (f Float) IsZero() bool {
	return float64(f) == 0.0
}

// IsNaN implements num.Num.
func (f Float) IsNaN() bool {
	return math.IsNaN(float64(f))
}

// Abs implements num.Num.
func (f Float) Abs() Float {
	return Float(math.Abs(float64(f)))
}

func (f Float) Exp(n Float) Float {
	return n.NewF(math.Exp(n.Float()))
}

func (f Float) Sub(n Float) Float {
	return Float(float64(f) - float64(n))
}

func (f Float) Add(n Float) Float {
	return Float(float64(f) + float64(n))
}

func (f Float) Div(n Float) Float {
	return Float(float64(f) / float64(n))
}

func (f Float) Mul(n Float) Float {
	return Float(float64(f) * float64(n))
}

func (f Float) LessThan(n Float) bool {
	return f < n
}

func (f Float) GreaterThan(n Float) bool {
	return f > n
}

func (f Float) LessThanOrEqual(n Float) bool {
	return f <= n
}

func (f Float) GreaterThanOrEqual(n Float) bool {
	return f >= n
}

func (f Float) Equal(n Float) bool {
	return f == n
}

func (f Float) Sign() int {
	if f < 0 {
		return -1
	} else if f > 0 {
		return 1
	}
	return 0
}

func Neg(n Float) Float {
	return -n
}

func (f Float) String() string {
	return fmt.Sprint(float64(f))
}
