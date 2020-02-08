package fluffy

import (
	"math"
)

type Term struct {
	Name           string
	MembershipFunc interface{ Value(x float64) float64 }
}

func (t *Term) MembershipValue(x float64) float64 {
	if t.MembershipFunc != nil {
		return t.MembershipFunc.Value(x)
	}
	return math.NaN()
}
