package mf

import (
	"math"

	"github.com/andviro/fluffy/v2/num"
)

type Sigmoid[T num.Num[T]] struct {
	A, C T
}

type sigmoidDTO struct {
	Type string
	A    string `yaml:"a"`
	C    string `yaml:"c"`
}

func (f Sigmoid[T]) MarshalYAML() (any, error) {
	return sigmoidDTO{Type: "Sigmoid", C: f.C.String(), A: f.A.String()}, nil
}

func sigmoid[T num.Num[T]](x, a, c T) T {
	return num.NewF[T](1.0 / (1.0 + math.Exp(-a.Float()*(x.Float()-c.Float()))))
}

func (f *Sigmoid[T]) Value(x T) T {
	return sigmoid(x, f.A, f.C)
}

type DSigmoid[T num.Num[T]] struct {
	A1, C1, A2, C2 T
}

type dsigmoidDTO struct {
	Type string
	A1   string `yaml:"a1"`
	C1   string `yaml:"c1"`
	A2   string `yaml:"a2"`
	C2   string `yaml:"c2"`
}

func (f DSigmoid[T]) MarshalYAML() (any, error) {
	return dsigmoidDTO{Type: "DSigmoid", C1: f.C1.String(), A1: f.A1.String(), C2: f.C2.String(), A2: f.A2.String()}, nil
}

func (f *DSigmoid[T]) Value(x T) T {
	return sigmoid(x, f.A1, f.C1).Sub(sigmoid(x, f.A2, f.C2))
}
