package mf

import (
	"github.com/andviro/fluffy/num"
)

type Sigmoid struct {
	A, C num.Num
}

type sigmoidDTO struct {
	Type string
	A    num.Num `yaml:"a"`
	C    num.Num `yaml:"c"`
}

func sigmoid(x, a, c num.Num) num.Num {
	t := num.Neg(a.Mul(x.Sub(c)))
	t = num.Exp(t)
	return one.Div(one.Add(t))
}

func (f Sigmoid) MarshalYAML() (interface{}, error) {
	return sigmoidDTO{Type: "Sigmoid", C: f.C, A: f.A}, nil
}

func (f *Sigmoid) Value(x num.Num) num.Num {
	return sigmoid(x, f.A, f.C)
}

type DSigmoid struct {
	A1, C1, A2, C2 num.Num
}

type dsigmoidDTO struct {
	Type string
	A1   num.Num `yaml:"a1"`
	C1   num.Num `yaml:"c1"`
	A2   num.Num `yaml:"a2"`
	C2   num.Num `yaml:"c2"`
}

func (f DSigmoid) MarshalYAML() (interface{}, error) {
	return dsigmoidDTO{Type: "DSigmoid", C1: f.C1, A1: f.A1, C2: f.C2, A2: f.A2}, nil
}

func (f *DSigmoid) Value(x num.Num) num.Num {
	return sigmoid(x, f.A1, f.C1).Sub(sigmoid(x, f.A2, f.C2))
}
