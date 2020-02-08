package mf

import "math"

type Gaussian struct {
	C     float64
	Sigma float64
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

type RightGaussian Gaussian

func (f RightGaussian) Value(x float64) float64 {
	if x >= f.C {
		return 1.0
	}
	return math.Exp(-math.Pow(x-f.C, 2) / (2 * math.Pow(f.Sigma, 2)))
}
