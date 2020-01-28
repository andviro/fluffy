package mf

type Triangle struct {
	A, B, C float64
}

func (f *Triangle) Value(x float64) float64 {
	switch {
	case x == f.B:
		return 1.0
	case x <= f.A:
		return 0.0
	case x >= f.C:
		return 0.0
	case x < f.B:
		return (x - f.A) / (f.B - f.A)
	}
	return (f.C - x) / (f.C - f.B)
}
