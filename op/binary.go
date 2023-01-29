package op

import (
	"reflect"
	"runtime"

	"github.com/andviro/fluffy/num"
)

type Binary func(num.Num, num.Num) num.Num

func (b Binary) MarshalYAML() (interface{}, error) {
	return runtime.FuncForPC(reflect.ValueOf(b).Pointer()).Name(), nil
}

func (b Binary) IsZero() bool {
	return b == nil
}

var (
	one = num.NewI(1, 0)
	two = num.NewI(2, 0)
)

func Nilmax(a num.Num, b num.Num) num.Num {
	if a.Add(b).LessThan(one) {
		return num.Max(a, b)
	}
	return one
}

func Hsum(a num.Num, b num.Num) num.Num {
	// return (a + b - (2 * a * b)) / (1 - (a * b))
	return (a.Add(b).Sub(two.Mul(a).Mul(b))).Div(one.Sub(a.Mul(b)))
}

func Esum(a num.Num, b num.Num) num.Num {
	// return (a + b) / (1 + a*b)
	return a.Add(b).Div(one.Add(a.Mul(b)))
}

func Drs(a num.Num, b num.Num) num.Num {
	switch {
	case a.IsZero():
		return b
	case b.IsZero():
		return a
	}
	return one
}

func Bsum(a num.Num, b num.Num) num.Num {
	return num.Min(one, a.Add(b))
}

func Probor(a num.Num, b num.Num) num.Num {
	// return a + b - (a * b)
	return a.Add(b).Sub(a.Mul(b))
}

func Mul(a num.Num, b num.Num) num.Num {
	return a.Mul(b)
}

var Max = num.Max

var Min = num.Min
