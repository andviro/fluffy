package mf

import (
	"github.com/andviro/fluffy/v2/num"
)

type LeftLinear struct {
	A, B num.Num
}

type linearDTO struct {
	Type string  `yaml:"type"`
	A    num.Num `yaml:"a"`
	B    num.Num `yaml:"b"`
}

func (l LeftLinear) MarshalYAML() (interface{}, error) {
	return linearDTO{Type: "LeftLinear", A: l.A, B: l.B}, nil
}

func (f LeftLinear) Value(x num.Num) num.Num {
	switch {
	case x.LessThanOrEqual(f.A):
		return one
	case x.GreaterThanOrEqual(f.B):
		return num.ZERO
	}
	return f.B.Sub(x).Div(f.B.Sub(f.A))
}

type RightLinear LeftLinear

func (l RightLinear) MarshalYAML() (interface{}, error) {
	return linearDTO{Type: "RightLinear", A: l.A, B: l.B}, nil
}

func (f RightLinear) Value(x num.Num) num.Num {
	switch {
	case x.LessThanOrEqual(f.A):
		return num.ZERO
	case x.GreaterThanOrEqual(f.B):
		return one
	}
	return x.Sub(f.A).Div(f.B.Sub(f.A))
}
