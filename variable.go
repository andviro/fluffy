package fluffy

import "github.com/shopspring/decimal"

type Variable struct {
	Name  VariableName    `yaml:"name"`
	Terms []Term          `yaml:"terms"`
	XMin  decimal.Decimal `yaml:"xMin"`
	XMax  decimal.Decimal `yaml:"xMax"`

	value      decimal.Decimal
	termValues map[TermName]decimal.Decimal
}

func (v *Variable) SetValue(value decimal.Decimal) {
	v.value = value
	v.termValues = make(map[TermName]decimal.Decimal)
	for _, t := range v.Terms {
		m := t.MembershipValue(value)
		v.termValues[t.Name] = m
	}
}

func (v *Variable) GetTermValue(term TermName) decimal.Decimal {
	return v.termValues[term]
}

func (v *Variable) GetTermValues() map[TermName]decimal.Decimal {
	vs := make(map[TermName]decimal.Decimal, len(v.termValues))
	for k, v := range v.termValues {
		vs[k] = v
	}
	return vs
}

func (v *Variable) SetTermValues(src map[TermName]decimal.Decimal) {
	v.termValues = make(map[TermName]decimal.Decimal, len(src))
	for k, val := range src {
		v.termValues[k] = val
	}
}

func (v *Variable) GetValue() decimal.Decimal {
	return v.value
}
