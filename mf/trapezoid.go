package mf

import "github.com/andviro/fluffy/num"

type Trapezoid struct {
	A, B, C, D num.Num
}

func (f Trapezoid) MarshalYAML() (interface{}, error) {
	return struct {
		Type string  `yaml:"type"`
		A    num.Num `yaml:"a"`
		B    num.Num `yaml:"b"`
		C    num.Num `yaml:"c"`
		D    num.Num `yaml:"d"`
	}{"Trapezoid", f.A, f.B, f.C, f.D}, nil
}

func (f *Trapezoid) Value(x num.Num) num.Num {
	switch {
	case x.GreaterThanOrEqual(f.B) && x.LessThanOrEqual(f.C):
		return one
	case x.LessThanOrEqual(f.A):
		return num.ZERO
	case x.GreaterThanOrEqual(f.D):
		return num.ZERO
	case x.LessThan(f.B):
		return x.Sub(f.A).Div(f.B.Sub(f.A))
	}
	return f.D.Sub(x).Div(f.D.Sub(f.C))
}
