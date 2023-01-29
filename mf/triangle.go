package mf

import "github.com/shopspring/decimal"

type Triangle struct {
	A, B, C decimal.Decimal
}

func (f Triangle) MarshalYAML() (interface{}, error) {
	return struct {
		Type string          `yaml:"type"`
		A    decimal.Decimal `yaml:"a"`
		B    decimal.Decimal `yaml:"b"`
		C    decimal.Decimal `yaml:"c"`
	}{"Triangle", f.A, f.B, f.C}, nil
}

func (f Triangle) Value(x decimal.Decimal) decimal.Decimal {
	switch {
	case x.Equal(f.B):
		return one
	case x.LessThan(f.A):
		return decimal.Zero
	case x.GreaterThanOrEqual(f.C):
		return decimal.Zero
	case x.LessThanOrEqual(f.B):
		return x.Sub(f.A).Div(f.B.Sub(f.A))
	}
	return f.C.Sub(x).Div(f.C.Sub(f.B))
}
