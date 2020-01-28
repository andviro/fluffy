package mf

type Singleton struct {
	A float64
}

func (f *Singleton) Value(x float64) float64 {
	if x == f.A {
		return 1.0
	}
	return 0.0
}
