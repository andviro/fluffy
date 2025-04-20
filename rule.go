package fluffy

import (
	"fmt"

	"github.com/andviro/fluffy/v2/num"
	"github.com/andviro/fluffy/v2/op"
)

var Epsilon = 0.0001

type Antecedent[T num.Num[T]] interface {
	Evaluator[T]
	MarshalYAML() (interface{}, error)
	Valid(func(variable VariableName, term TermName) error) error
}

type Evaluator[T num.Num[T]] interface {
	Evaluate(fis FIS[T]) T
}

type Rule[T num.Num[T]] struct {
	Weight      T
	AndMethod   op.Binary[T]
	OrMethod    op.Binary[T]
	Antecedent  Antecedent[T]
	Consequents []Clause[T]
}

func (r Rule[T]) MarshalYAML() (interface{}, error) {
	return struct {
		Weight      float64
		AndMethod   op.Binary[T]  `yaml:"andMethod,omitempty"`
		OrMethod    op.Binary[T]  `yaml:"orMethod,omitempty"`
		Antecedent  Antecedent[T] `yaml:"antecedent"`
		Consequents []Clause[T]   `yaml:"consequents"`
	}{
		Weight:      r.Weight.Float(),
		AndMethod:   r.AndMethod,
		OrMethod:    r.OrMethod,
		Antecedent:  r.Antecedent,
		Consequents: r.Consequents,
	}, nil
}

type rule[T num.Num[T]] struct {
	*Rule[T]
	FIS[T]
}

func (r rule[T]) And(a T, b T) T {
	if r.AndMethod != nil {
		return r.AndMethod(a, b)
	}
	return r.FIS.And(a, b)
}

func (r rule[T]) Or(a T, b T) T {
	if r.OrMethod != nil {
		return r.OrMethod(a, b)
	}
	return r.FIS.Or(a, b)
}

func (r *Rule[T]) Evaluate(fis FIS[T]) {
	w := r.Antecedent.Evaluate(rule[T]{Rule: r, FIS: fis})
	if !w.IsZero() {
		for _, c := range r.Consequents {
			fis.Activate(c, w.Mul(r.Weight))
		}
	}
}

func (r Rule[T]) String() string {
	return fmt.Sprintf("%s => %v", r.Antecedent, r.Consequents)
}
