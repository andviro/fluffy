package fluffy

import (
	"github.com/shopspring/decimal"
)

type MembershipFunc interface {
	Value(x decimal.Decimal) decimal.Decimal
	MarshalYAML() (interface{}, error)
}

type Term struct {
	Name           TermName       `yaml:"name"`
	MembershipFunc MembershipFunc `yaml:"mf"`
}

func (t *Term) MembershipValue(x decimal.Decimal) decimal.Decimal {
	if t.MembershipFunc == nil {
		return decimal.Zero
	}
	return t.MembershipFunc.Value(x)
}
