package mf

type LeftLinear struct {
	A, B float64
}

type linearDTO struct {
	Type string  `yaml:"type"`
	A    float64 `yaml:"a"`
	B    float64 `yaml:"b"`
}

func (l LeftLinear) MarshalYAML() (interface{}, error) {
	return linearDTO{Type: "LeftLinear", A: l.A, B: l.B}, nil
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

func (l RightLinear) MarshalYAML() (interface{}, error) {
	return linearDTO{Type: "RightLinear", A: l.A, B: l.B}, nil
}

func (f RightLinear) Value(x float64) float64 {
	switch {
	case x <= f.A:
		return 0.0
	case x >= f.B:
		return 1.0
	}
	return (x - f.A) / (f.B - f.A)
}
