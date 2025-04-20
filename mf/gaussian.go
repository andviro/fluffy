package mf

import (
	"github.com/andviro/fluffy/v2/num"
)

type Gaussian[T num.Num[T]] struct {
	C     T
	Sigma T
}

type gaussianDTO struct {
	Type  string `yaml:"type"`
	C     string `yaml:"c"`
	Sigma string `yaml:"sigma"`
}

func (f Gaussian[T]) MarshalYAML() (any, error) {
	return gaussianDTO{Type: "Gaussian", C: f.C.String(), Sigma: f.Sigma.String()}, nil
}

func (f Gaussian[T]) Value(x T) T {
	return x.Exp((num.Sqr(x.Sub(f.C)).Div(num.Two[T]().Mul(num.Sqr(f.Sigma)))).Neg())
}

type LeftGaussian[T num.Num[T]] Gaussian[T]

func (f LeftGaussian[T]) Value(x T) T {
	if x.LessThanOrEqual(f.C) {
		return num.One[T]()
	}
	return Gaussian[T](f).Value(x)
}

func (f LeftGaussian[T]) MarshalYAML() (any, error) {
	return gaussianDTO{Type: "LeftGaussian", C: f.C.String(), Sigma: f.Sigma.String()}, nil
}

type RightGaussian[T num.Num[T]] Gaussian[T]

func (f RightGaussian[T]) Value(x T) T {
	if x.GreaterThanOrEqual(f.C) {
		return num.One[T]()
	}
	return Gaussian[T](f).Value(x)
}

func (f RightGaussian[T]) MarshalYAML() (any, error) {
	return gaussianDTO{Type: "RightGaussian", C: f.C.String(), Sigma: f.Sigma.String()}, nil
}
