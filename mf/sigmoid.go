package mf

import "math"

type Sigmoid struct {
	A, C float64
}

type sigmoidDTO struct {
	Type string
	A    float64 `yaml:"a"`
	C    float64 `yaml:"c"`
}

func (f Sigmoid) MarshalYAML() (interface{}, error) {
	return sigmoidDTO{Type: "Sigmoid", C: f.C, A: f.A}, nil
}

func (f *Sigmoid) Value(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-f.A*(x-f.C)))
}

type DSigmoid struct {
	A1, C1, A2, C2 float64
}

type dsigmoidDTO struct {
	Type string
	A1   float64 `yaml:"a1"`
	C1   float64 `yaml:"c1"`
	A2   float64 `yaml:"a2"`
	C2   float64 `yaml:"c2"`
}

func (f DSigmoid) MarshalYAML() (interface{}, error) {
	return dsigmoidDTO{Type: "DSigmoid", C1: f.C1, A1: f.A1, C2: f.C2, A2: f.A2}, nil
}

func (f *DSigmoid) Value(x float64) float64 {
	return 1.0/(1.0+math.Exp(-f.A1*(x-f.C1))) - 1.0/(1.0+math.Exp(-f.A2*(x-f.C2)))
}
