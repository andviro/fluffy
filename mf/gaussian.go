package mf

import "math"

type Gaussian struct {
	C     float64
	Sigma float64
}

type gaussianDTO struct {
	Type  string  `yaml:"type"`
	C     float64 `yaml:"c"`
	Sigma float64 `yaml:"sigma"`
}

func (f Gaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "Gaussian", C: f.C, Sigma: f.Sigma}, nil
}

func (f Gaussian) Value(x float64) float64 {
	return math.Exp(-math.Pow(x-f.C, 2) / (2 * math.Pow(f.Sigma, 2)))
}

type LeftGaussian Gaussian

func (f LeftGaussian) Value(x float64) float64 {
	if x <= f.C {
		return 1.0
	}
	return math.Exp(-math.Pow(x-f.C, 2) / (2 * math.Pow(f.Sigma, 2)))
}

func (f LeftGaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "LeftGaussian", C: f.C, Sigma: f.Sigma}, nil
}

type RightGaussian Gaussian

func (f RightGaussian) Value(x float64) float64 {
	if x >= f.C {
		return 1.0
	}
	return math.Exp(-math.Pow(x-f.C, 2) / (2 * math.Pow(f.Sigma, 2)))
}

func (f RightGaussian) MarshalYAML() (interface{}, error) {
	return gaussianDTO{Type: "RightGaussian", C: f.C, Sigma: f.Sigma}, nil
}
