package fluffy

type Variable struct {
	Name       string
	Terms      []Term
	value      float64
	termValues map[string]float64
}

func (v *Variable) SetValue(value float64) {
	v.value = value
	v.termValues = make(map[string]float64)
	for _, t := range v.Terms {
		v.termValues[t.Name] = t.MembershipValue(value)
	}
}
