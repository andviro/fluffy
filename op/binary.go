package op

import (
	"reflect"
	"runtime"

	"github.com/andviro/fluffy/v2/num"
)

type Binary[T num.Num[T]] func(T, T) T

func (b Binary[T]) MarshalYAML() (interface{}, error) {
	return runtime.FuncForPC(reflect.ValueOf(b).Pointer()).Name(), nil
}

func (b Binary[T]) IsZero() bool {
	return b == nil
}

func one[T num.Num[T]]() (res T) {
	return res.NewI(1, 0)
}

func two[T num.Num[T]]() (res T) {
	return res.NewI(2, 0)
}

func Nilmax[T num.Num[T]](a T, b T) T {
	if a.Add(b).LessThan(one[T]()) {
		return num.Max(a, b)
	}
	return one[T]()
}

func Hsum[T num.Num[T]](a T, b T) T {
	// return (a + b - (2 * a * b)) / (1 - (a * b))
	return (a.Add(b).Sub(two[T]().Mul(a).Mul(b))).Div(one[T]().Sub(a.Mul(b)))
}

func Esum[T num.Num[T]](a T, b T) T {
	// return (a + b) / (1 + a*b)
	return a.Add(b).Div(one[T]().Add(a.Mul(b)))
}

func Drs[T num.Num[T]](a T, b T) T {
	switch {
	case a.IsZero():
		return b
	case b.IsZero():
		return a
	}
	return one[T]()
}

func Bsum[T num.Num[T]](a T, b T) T {
	return num.Min(one[T](), a.Add(b))
}

func Probor[T num.Num[T]](a T, b T) T {
	// return a + b - (a * b)
	return a.Add(b).Sub(a.Mul(b))
}

func Mul[T num.Num[T]](a T, b T) T {
	return a.Mul(b)
}

func Max[T num.Num[T]](a T, rest ...T) T {
	return num.Max(a, rest...)
}

func Min[T num.Num[T]](a T, rest ...T) T {
	return num.Min(a, rest...)
}
