package fluffy

import (
	"fmt"
	"strings"

	"github.com/andviro/fluffy/v2/num"
)

type VariableName string

type TermName string

type Clause[T num.Num[T]] struct {
	Variable VariableName `yaml:"variable" parser:"@Ident"`
	Term     TermName     `yaml:"term" parser:"'=' @Ident"`
}

func (c Clause[T]) Valid(f func(VariableName, TermName) error) error {
	return f(c.Variable, c.Term)
}

func (c Clause[T]) MarshalYAML() (interface{}, error) {
	return map[interface{}]interface{}{
		c.Variable: c.Term,
	}, nil
}

func C[T num.Num[T]](variable VariableName, term TermName) Clause[T] {
	return Clause[T]{variable, term}
}

func (c Clause[T]) String() string {
	return fmt.Sprintf("%s=%s", c.Variable, c.Term)
}

func (c Clause[T]) Evaluate(fis FIS[T]) T {
	v := fis.GetInput(c.Variable)
	return v.GetTermValue(c.Term)
}

type Connector[T num.Num[T]] []Antecedent[T]

func (a Connector[T]) Valid(f func(VariableName, TermName) error) error {
	for _, clause := range a {
		if err := clause.Valid(f); err != nil {
			return err
		}
	}
	return nil
}

type And[T num.Num[T]] Connector[T]

func (a And[T]) Valid(f func(VariableName, TermName) error) error {
	return Connector[T](a).Valid(f)
}

func (a And[T]) MarshalYAML() (interface{}, error) {
	return struct {
		And []Antecedent[T] `yaml:"and"`
	}{a}, nil
}

func (c Connector[T]) string(symbol string) string {
	var res []string
	for _, e := range c {
		res = append(res, fmt.Sprintf("%s", e))
	}
	return fmt.Sprintf("(%s)", strings.Join(res, symbol))
}

func (a And[T]) Evaluate(fis FIS[T]) T {
	if len(a) == 0 {
		return num.NaN[T]()
	}
	res := a[0].Evaluate(fis)
	for _, b := range a[1:] {
		res = fis.And(res, b.Evaluate(fis))
	}
	return res
}

func (a And[T]) String() string {
	return Connector[T](a).string(" and ")
}

type Or[T num.Num[T]] Connector[T]

func (a Or[T]) Valid(f func(VariableName, TermName) error) error {
	return Connector[T](a).Valid(f)
}

func (a Or[T]) MarshalYAML() (interface{}, error) {
	return struct {
		Or []Antecedent[T] `yaml:"or"`
	}{a}, nil
}

func (a Or[T]) Evaluate(fis FIS[T]) T {
	if len(a) == 0 {
		return num.ZERO[T]()
	}
	res := a[0].Evaluate(fis)
	for _, b := range a[1:] {
		res = fis.Or(res, b.Evaluate(fis))
	}
	return res
}

func (a Or[T]) String() string {
	return Connector[T](a).string(" or ")
}

type Not[T num.Num[T]] struct {
	Antecedent[T]
}

func (a Not[T]) Evaluate(fis FIS[T]) T {
	return num.One[T]().Sub(a.Antecedent.Evaluate(fis))
}

func (a Not[T]) String() string {
	return fmt.Sprintf("not %s", a.Antecedent)
}
