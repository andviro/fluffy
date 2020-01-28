package mf

type Trapezoid struct {
	A, B, C, D float64
}

func (f *Trapezoid) Value(x float64) float64 {
	switch {
	case x >= f.B && x <= f.C:
		return 1.0
	case x <= f.A:
		return 0.0
	case x >= f.D:
		return 0.0
	case x < f.B:
		return (x - f.A) / (f.B - f.A)
	}
	return (f.D - x) / (f.D - f.C)
}
