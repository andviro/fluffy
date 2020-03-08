package fluffy

type Variable struct {
	Name  VariableName
	Terms []Term

	value      float64
	termValues map[TermName]float64
}

func (v *Variable) SetValue(value float64) {
	v.value = value
	v.termValues = make(map[TermName]float64)
	for _, t := range v.Terms {
		m := t.MembershipValue(value)
		v.termValues[t.Name] = m
	}
}

func (v *Variable) GetTermValue(term TermName) float64 {
	return v.termValues[term]
}

func (v *Variable) GetValue() float64 {
	return v.value
}
