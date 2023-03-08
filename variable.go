package fluffy

import "github.com/andviro/fluffy/v2/num"

type Variable struct {
	Name  VariableName `yaml:"name"`
	Terms []Term       `yaml:"terms"`
	XMin  num.Num      `yaml:"xMin"`
	XMax  num.Num      `yaml:"xMax"`

	value      num.Num
	termValues map[TermName]num.Num
}

func (v *Variable) SetValue(value num.Num) {
	v.value = value
	v.termValues = make(map[TermName]num.Num)
	for _, t := range v.Terms {
		m := t.MembershipValue(value)
		v.termValues[t.Name] = m
	}
}

func (v *Variable) GetTermValue(term TermName) num.Num {
	return v.termValues[term]
}

func (v *Variable) GetTermValues() map[TermName]num.Num {
	vs := make(map[TermName]num.Num, len(v.termValues))
	for k, v := range v.termValues {
		vs[k] = v
	}
	return vs
}

func (v *Variable) SetTermValues(src map[TermName]num.Num) {
	v.termValues = make(map[TermName]num.Num, len(src))
	for k, val := range src {
		v.termValues[k] = val
	}
}

func (v *Variable) GetValue() num.Num {
	return v.value
}
