package mf

import "github.com/shopspring/decimal"

type Singleton struct {
	A decimal.Decimal
}

func (f Singleton) MarshalYAML() (interface{}, error) {
	return struct {
		Type string          `yaml:"type"`
		A    decimal.Decimal `yaml:"a"`
	}{"Singleton", f.A}, nil
}

func (f *Singleton) Value(x decimal.Decimal) decimal.Decimal {
	if x.Equal(f.A) {
		return one
	}
	return decimal.Zero
}
