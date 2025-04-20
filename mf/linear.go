package mf

import (
	"github.com/andviro/fluffy/v2/num"
)

type LeftLinear[T num.Num[T]] struct {
	A, B T
}

type linearDTO struct {
	Type string `yaml:"type"`
	A    string `yaml:"a"`
	B    string `yaml:"b"`
}

func (l LeftLinear[T]) MarshalYAML() (any, error) {
	return linearDTO{Type: "LeftLinear", A: l.A.String(), B: l.B.String()}, nil
}

func (f LeftLinear[T]) Value(x T) T {
	switch {
	case x.LessThanOrEqual(f.A):
		return num.One[T]()
	case x.GreaterThanOrEqual(f.B):
		return num.ZERO[T]()
	}
	return f.B.Sub(x).Div(f.B.Sub(f.A))
}

type RightLinear[T num.Num[T]] LeftLinear[T]

func (l RightLinear[T]) MarshalYAML() (any, error) {
	return linearDTO{Type: "RightLinear", A: l.A.String(), B: l.B.String()}, nil
}

func (f RightLinear[T]) Value(x T) T {
	switch {
	case x.LessThanOrEqual(f.A):
		return num.ZERO[T]()
	case x.GreaterThanOrEqual(f.B):
		return num.One[T]()
	}
	return x.Sub(f.A).Div(f.B.Sub(f.A))
}
