package mf

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var (
	one = decimal.NewFromInt(1)
	two = decimal.NewFromInt(2)
)

type Gaussian struct {
	C     decimal.Decimal
	Sigma decimal.Decimal
}

type gaussianDTO struct {
	Type  string          `yaml:"type"`
	C     decimal.Decimal `yaml:"c"`
	Sigma decimal.Decimal `yaml:"sigma"`
}

func (f Gaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "Gaussian", C: f.C, Sigma: f.Sigma}, nil
}

func (f Gaussian) Value(x decimal.Decimal) decimal.Decimal {
	// return math.Exp(-math.Pow(x-f.C, 2) / (2 * math.Pow(f.Sigma, 2)))
	fmt.Println("###", x, f.C, f.Sigma)
	t := x.Sub(f.C)
	t = t.Mul(t).Neg()
	t, _ = t.ExpHullAbrham(20)
	t = t.Div(f.Sigma.Mul(f.Sigma).Mul(two))
	return t
}

type LeftGaussian Gaussian

func (f LeftGaussian) Value(x decimal.Decimal) decimal.Decimal {
	if x.LessThanOrEqual(f.C) {
		return one
	}
	// return math.Exp(-math.Pow(x-f.C, 2) / (2 * math.Pow(f.Sigma, 2)))
	return Gaussian(f).Value(x)
}

func (f LeftGaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "LeftGaussian", C: f.C, Sigma: f.Sigma}, nil
}

type RightGaussian Gaussian

func (f RightGaussian) Value(x decimal.Decimal) decimal.Decimal {
	if x.GreaterThanOrEqual(f.C) {
		return one
	}
	return Gaussian(f).Value(x)
}

func (f RightGaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "RightGaussian", C: f.C, Sigma: f.Sigma}, nil
}
