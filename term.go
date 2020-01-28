package fluffy

type MembershipFunc interface{ Value(x float64) float64 }

type Term struct {
	Name           string
	MembershipFunc MembershipFunc
}

func (t *Term) MembershipValue(x float64) float64 {
	if t.MembershipFunc != nil {
		return t.MembershipFunc.Value(x)
	}
	return 0
}
