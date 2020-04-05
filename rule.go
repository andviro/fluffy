package fluffy

import (
	"fmt"

	"github.com/andviro/fluffy/op"
)

type Evaluator interface {
	Evaluate(fis FIS) float64
}

type Rule struct {
	Weight      float64
	AndMethod   op.Binary
	OrMethod    op.Binary
	Antecedent  Evaluator
	Consequents []Clause
}

type rule struct {
	*Rule
	FIS
}

func (r rule) And(a float64, b float64) float64 {
	if r.AndMethod != nil {
		return r.AndMethod(a, b)
	}
	return r.FIS.And(a, b)
}

func (r rule) Or(a float64, b float64) float64 {
	if r.OrMethod != nil {
		return r.OrMethod(a, b)
	}
	return r.FIS.Or(a, b)
}

func (r *Rule) Evaluate(fis FIS) {
	w := r.Antecedent.Evaluate(rule{Rule: r, FIS: fis})
	for _, c := range r.Consequents {
		fis.Activate(c, w*r.Weight)
	}
}

func (r Rule) String() string {
	return fmt.Sprintf("%s => %v", r.Antecedent, r.Consequents)
}
