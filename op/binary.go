package op

import (
	"math"
	"reflect"
	"runtime"
)

type Binary func(float64, float64) float64

func (b Binary) MarshalYAML() (interface{}, error) {
	return runtime.FuncForPC(reflect.ValueOf(b).Pointer()).Name(), nil
}

func (b Binary) IsZero() bool {
	return b == nil
}

func Nilmax(a float64, b float64) float64 {
	if a+b < 1 {
		return math.Max(a, b)
	}
	return 1
}

func Hsum(a float64, b float64) float64 {
	return (a + b - (2 * a * b)) / (1 - (a * b))
}

func Esum(a float64, b float64) float64 {
	return (a + b) / (1 + a*b)
}

func Drs(a float64, b float64) float64 {
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	}
	return 1
}

func Bsum(a float64, b float64) float64 {
	return math.Min(1, a+b)
}

func Probor(a float64, b float64) float64 {
	return a + b - (a * b)
}

func Mul(a float64, b float64) float64 {
	return a * b
}

var Max = math.Max

var Min = math.Min
