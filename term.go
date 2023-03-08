package fluffy

import (
	"github.com/andviro/fluffy/v2/num"
)

type MembershipFunc interface {
	Value(x num.Num) num.Num
	MarshalYAML() (interface{}, error)
}

type Term struct {
	Name           TermName       `yaml:"name"`
	MembershipFunc MembershipFunc `yaml:"mf"`
}

func (t *Term) MembershipValue(x num.Num) num.Num {
	if t.MembershipFunc == nil {
		return num.NaN
	}
	return t.MembershipFunc.Value(x)
}
