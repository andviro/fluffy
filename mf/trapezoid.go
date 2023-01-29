package mf

import "github.com/shopspring/decimal"

type Trapezoid struct {
	A, B, C, D decimal.Decimal
}

func (f Trapezoid) MarshalYAML() (interface{}, error) {
	return struct {
		Type string          `yaml:"type"`
		A    decimal.Decimal `yaml:"a"`
		B    decimal.Decimal `yaml:"b"`
		C    decimal.Decimal `yaml:"c"`
		D    decimal.Decimal `yaml:"d"`
	}{"Trapezoid", f.A, f.B, f.C, f.D}, nil
}

func (f *Trapezoid) Value(x decimal.Decimal) decimal.Decimal {
	switch {
	case x.GreaterThanOrEqual(f.B) && x.LessThanOrEqual(f.C):
		return one
	case x.LessThanOrEqual(f.A):
		return decimal.Zero
	case x.GreaterThanOrEqual(f.D):
		return decimal.Zero
	case x.LessThan(f.B):
		return x.Sub(f.A).Div(f.B.Sub(f.A))
	}
	return f.D.Sub(x).Div(f.D.Sub(f.C))
}
