package fluffy

type Variable struct {
	Name  VariableName `yaml:"name"`
	Terms []Term       `yaml:"terms"`
	XMin  float64      `yaml:"xMin"`
	XMax  float64      `yaml:"xMax"`

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

func (v *Variable) GetTermValues() map[TermName]float64 {
	vs := make(map[TermName]float64, len(v.termValues))
	for k, v := range v.termValues {
		vs[k] = v
	}
	return vs
}

func (v *Variable) SetTermValues(src map[TermName]float64) {
	v.termValues = make(map[TermName]float64, len(src))
	for k, val := range src {
		v.termValues[k] = val
	}
}

func (v *Variable) GetValue() float64 {
	return v.value
}
