package mf

import (
	"github.com/shopspring/decimal"
)

type Sigmoid struct {
	A, C decimal.Decimal
}

type sigmoidDTO struct {
	Type string
	A    decimal.Decimal `yaml:"a"`
	C    decimal.Decimal `yaml:"c"`
}

func sigmoid(x, a, c decimal.Decimal) decimal.Decimal {
	t := a.Mul(x.Sub(c)).Neg()
	t, _ = t.ExpHullAbrham(20)
	return one.Div(one.Add(t))
}

func (f Sigmoid) MarshalYAML() (interface{}, error) {
	return sigmoidDTO{Type: "Sigmoid", C: f.C, A: f.A}, nil
}

func (f *Sigmoid) Value(x decimal.Decimal) decimal.Decimal {
	return sigmoid(x, f.A, f.C)
}

type DSigmoid struct {
	A1, C1, A2, C2 decimal.Decimal
}

type dsigmoidDTO struct {
	Type string
	A1   decimal.Decimal `yaml:"a1"`
	C1   decimal.Decimal `yaml:"c1"`
	A2   decimal.Decimal `yaml:"a2"`
	C2   decimal.Decimal `yaml:"c2"`
}

func (f DSigmoid) MarshalYAML() (interface{}, error) {
	return dsigmoidDTO{Type: "DSigmoid", C1: f.C1, A1: f.A1, C2: f.C2, A2: f.A2}, nil
}

func (f *DSigmoid) Value(x decimal.Decimal) decimal.Decimal {
	return sigmoid(x, f.A1, f.C1).Sub(sigmoid(x, f.A2, f.C2))
}
