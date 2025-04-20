package mf

import "github.com/andviro/fluffy/v2/num"

type Trapezoid[T num.Num[T]] struct {
	A, B, C, D T
}

func (f Trapezoid[T]) MarshalYAML() (any, error) {
	return struct {
		Type string `yaml:"type"`
		A    string `yaml:"a"`
		B    string `yaml:"b"`
		C    string `yaml:"c"`
		D    string `yaml:"d"`
	}{"Trapezoid", f.A.String(), f.B.String(), f.C.String(), f.D.String()}, nil
}

func (f *Trapezoid[T]) Value(x T) T {
	switch {
	case x.GreaterThanOrEqual(f.B) && x.LessThanOrEqual(f.C):
		return num.One[T]()
	case x.LessThanOrEqual(f.A):
		return num.ZERO[T]()
	case x.GreaterThanOrEqual(f.D):
		return num.ZERO[T]()
	case x.LessThan(f.B):
		return x.Sub(f.A).Div(f.B.Sub(f.A))
	}
	return f.D.Sub(x).Div(f.D.Sub(f.C))
}
