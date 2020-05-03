package fluffy

import (
	"math"
)

type MembershipFunc interface {
	Value(x float64) float64
	MarshalYAML() (interface{}, error)
}

type Term struct {
	Name           TermName       `yaml:"name"`
	MembershipFunc MembershipFunc `yaml:"mf"`
}

func (t *Term) MembershipValue(x float64) float64 {
	if t.MembershipFunc != nil {
		return t.MembershipFunc.Value(x)
	}
	return math.NaN()
}
