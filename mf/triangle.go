package mf

import (
	"github.com/andviro/fluffy/num"
)

type Triangle struct {
	A, B, C num.Num
}

func (f Triangle) MarshalYAML() (interface{}, error) {
	return struct {
		Type string  `yaml:"type"`
		A    num.Num `yaml:"a"`
		B    num.Num `yaml:"b"`
		C    num.Num `yaml:"c"`
	}{"Triangle", f.A, f.B, f.C}, nil
}

func (f Triangle) Value(x num.Num) num.Num {
	switch {
	case x.Equal(f.B):
		return one
	case x.LessThan(f.A):
		return num.ZERO
	case x.GreaterThanOrEqual(f.C):
		return num.ZERO
	case x.LessThanOrEqual(f.B):
		return x.Sub(f.A).Div(f.B.Sub(f.A))
	}
	return f.C.Sub(x).Div(f.C.Sub(f.B))
}
