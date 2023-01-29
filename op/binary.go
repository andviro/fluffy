package op

import (
	"reflect"
	"runtime"

	"github.com/shopspring/decimal"
)

type Binary func(decimal.Decimal, decimal.Decimal) decimal.Decimal

func (b Binary) MarshalYAML() (interface{}, error) {
	return runtime.FuncForPC(reflect.ValueOf(b).Pointer()).Name(), nil
}

func (b Binary) IsZero() bool {
	return b == nil
}

var (
	one = decimal.NewFromInt(1)
	two = decimal.NewFromInt(2)
)

func Nilmax(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	if a.Add(b).LessThan(one) {
		return decimal.Max(a, b)
	}
	return one
}

func Hsum(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	// return (a + b - (2 * a * b)) / (1 - (a * b))
	return (a.Add(b).Sub(two.Mul(a).Mul(b))).Div(one.Sub(a.Mul(b)))
}

func Esum(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	// return (a + b) / (1 + a*b)
	return a.Add(b).Div(one.Add(a.Mul(b)))
}

func Drs(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	switch {
	case a.IsZero():
		return b
	case b.IsZero():
		return a
	}
	return one
}

func Bsum(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	return decimal.Min(one, a.Add(b))
}

func Probor(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	// return a + b - (a * b)
	return a.Add(b).Sub(a.Mul(b))
}

func Mul(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	return a.Mul(b)
}

var Max = decimal.Max

var Min = decimal.Min
