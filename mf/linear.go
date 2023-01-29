package mf

import "github.com/shopspring/decimal"

type LeftLinear struct {
	A, B decimal.Decimal
}

type linearDTO struct {
	Type string          `yaml:"type"`
	A    decimal.Decimal `yaml:"a"`
	B    decimal.Decimal `yaml:"b"`
}

func (l LeftLinear) MarshalYAML() (interface{}, error) {
	return linearDTO{Type: "LeftLinear", A: l.A, B: l.B}, nil
}

func (f LeftLinear) Value(x decimal.Decimal) decimal.Decimal {
	switch {
	case x.LessThanOrEqual(f.A):
		return one
	case x.GreaterThanOrEqual(f.B):
		return decimal.Zero
	}
	return f.B.Sub(x).Div(f.B.Sub(f.A))
}

type RightLinear LeftLinear

func (l RightLinear) MarshalYAML() (interface{}, error) {
	return linearDTO{Type: "RightLinear", A: l.A, B: l.B}, nil
}

func (f RightLinear) Value(x decimal.Decimal) decimal.Decimal {
	switch {
	case x.LessThanOrEqual(f.A):
		return decimal.Zero
	case x.GreaterThanOrEqual(f.B):
		return one
	}
	return x.Sub(f.A).Div(f.B.Sub(f.A))
}
