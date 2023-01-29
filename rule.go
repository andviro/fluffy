package fluffy

import (
	"fmt"

	"github.com/andviro/fluffy/op"
	"github.com/shopspring/decimal"
)

var Epsilon = 0.0001

type Antecedent interface {
	Evaluator
	MarshalYAML() (interface{}, error)
}

type Evaluator interface {
	Evaluate(fis FIS) decimal.Decimal
}

type Rule struct {
	Weight      decimal.Decimal `yaml:"weight"`
	AndMethod   op.Binary       `yaml:"andMethod,omitempty"`
	OrMethod    op.Binary       `yaml:"orMethod,omitempty"`
	Antecedent  Antecedent      `yaml:"antecedent"`
	Consequents []Clause        `yaml:"consequents"`
}

type rule struct {
	*Rule
	FIS
}

func (r rule) And(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	if r.AndMethod != nil {
		return r.AndMethod(a, b)
	}
	return r.FIS.And(a, b)
}

func (r rule) Or(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	if r.OrMethod != nil {
		return r.OrMethod(a, b)
	}
	return r.FIS.Or(a, b)
}

func (r *Rule) Evaluate(fis FIS) {
	w := r.Antecedent.Evaluate(rule{Rule: r, FIS: fis})
	if !w.IsZero() {
		for _, c := range r.Consequents {
			fis.Activate(c, w.Mul(r.Weight))
		}
	}
}

func (r Rule) String() string {
	return fmt.Sprintf("%s => %v", r.Antecedent, r.Consequents)
}
