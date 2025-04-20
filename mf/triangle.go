package mf

import (
	"github.com/andviro/fluffy/v2/num"
)

type Triangle[T num.Num[T]] struct {
	A, B, C T
}

func (f Triangle[T]) MarshalYAML() (any, error) {
	return struct {
		Type string `yaml:"type"`
		A    string `yaml:"a"`
		B    string `yaml:"b"`
		C    string `yaml:"c"`
	}{"Triangle", f.A.String(), f.B.String(), f.C.String()}, nil
}

func (f Triangle[T]) Value(x T) T {
	switch {
	case x.Equal(f.B):
		return num.One[T]()
	case x.LessThan(f.A):
		return num.ZERO[T]()
	case x.GreaterThanOrEqual(f.C):
		return num.ZERO[T]()
	case x.LessThanOrEqual(f.B):
		return x.Sub(f.A).Div(f.B.Sub(f.A))
	}
	return f.C.Sub(x).Div(f.C.Sub(f.B))
}
