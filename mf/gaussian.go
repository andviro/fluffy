package mf

import (
	"github.com/andviro/fluffy/num"
)

var (
	one = num.NewI(1, 0)
	two = num.NewI(2, 0)
)

type Gaussian struct {
	C     num.Num
	Sigma num.Num
}

type gaussianDTO struct {
	Type  string  `yaml:"type"`
	C     num.Num `yaml:"c"`
	Sigma num.Num `yaml:"sigma"`
}

func (f Gaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "Gaussian", C: f.C, Sigma: f.Sigma}, nil
}

func (f Gaussian) Value(x num.Num) num.Num {
	// return math.Exp(-math.Pow(x-f.C, 2) / (2 * math.Pow(f.Sigma, 2)))
	t := x.Sub(f.C)
	t = num.ZERO.Sub(t.Mul(t))
	t = num.Exp(t)
	t = t.Div(f.Sigma.Mul(f.Sigma).Mul(two))
	return t
}

type LeftGaussian Gaussian

func (f LeftGaussian) Value(x num.Num) num.Num {
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

func (f RightGaussian) Value(x num.Num) num.Num {
	if x.GreaterThanOrEqual(f.C) {
		return one
	}
	return Gaussian(f).Value(x)
}

func (f RightGaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "RightGaussian", C: f.C, Sigma: f.Sigma}, nil
}
