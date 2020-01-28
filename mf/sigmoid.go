package mf

import "math"

type Sigmoid struct {
	A, C float64
}

func (f *Sigmoid) Value(x float64) float64 {
	return 1.0 / math.Exp(-f.A*(x-f.C))
}

type DSigmoid struct {
	A1, C1, A2, C2 float64
}

func (f *DSigmoid) Value(x float64) float64 {
	return 1.0/math.Exp(-f.A1*(x-f.C1)) - 1.0/math.Exp(-f.A2*(x-f.C2))
}
