package fluffy

type Variable struct {
	Name  string
	Terms []Term

	value      float64
	termValues map[string]float64
}

func (v *Variable) SetValue(value float64) {
	v.value = value
	v.termValues = make(map[string]float64)
	for _, t := range v.Terms {
		m := t.MembershipValue(value)
		v.termValues[t.Name] = m
	}
}

func (v *Variable) GetTermValue(term string) float64 {
	return v.termValues[term]
}

func (v *Variable) GetValue() float64 {
	return v.value
}
