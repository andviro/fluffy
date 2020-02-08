package mf

type LeftLinear struct {
	A, B float64
}

func (f LeftLinear) Value(x float64) float64 {
	switch {
	case x <= f.A:
		return 1.0
	case x >= f.B:
		return 0.0
	}
	return (f.B - x) / (f.B - f.A)
}

type RightLinear LeftLinear

func (f RightLinear) Value(x float64) float64 {
	switch {
	case x <= f.A:
		return 0.0
	case x >= f.B:
		return 1.0
	}
	return (x - f.A) / (f.B - f.A)
}
