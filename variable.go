package fluffy

import "github.com/andviro/fluffy/v2/num"

type Variable[T num.Num[T]] struct {
	Name  VariableName `yaml:"name"`
	Terms []Term[T]    `yaml:"terms"`
	XMin  T            `yaml:"xMin"`
	XMax  T            `yaml:"xMax"`

	value      T
	termValues map[TermName]T
}

func (v *Variable[T]) SetValue(value T) {
	v.value = value
	v.termValues = make(map[TermName]T)
	for _, t := range v.Terms {
		m := t.MembershipValue(value)
		v.termValues[t.Name] = m
	}
}

func (v *Variable[T]) GetTermValue(term TermName) T {
	return v.termValues[term]
}

func (v *Variable[T]) GetTermValues() map[TermName]T {
	vs := make(map[TermName]T, len(v.termValues))
	for k, v := range v.termValues {
		vs[k] = v
	}
	return vs
}

func (v *Variable[T]) SetTermValues(src map[TermName]T) {
	v.termValues = make(map[TermName]T, len(src))
	for k, val := range src {
		v.termValues[k] = val
	}
}

func (v *Variable[T]) GetValue() T {
	return v.value
}
