package fluffy

import (
	"fmt"
	"strings"

	"github.com/andviro/fluffy/v2/num"
)

type VariableName string

type TermName string

type Clause struct {
	Variable VariableName `yaml:"variable" parser:"@Ident"`
	Term     TermName     `yaml:"term" parser:"'=' @Ident"`
}

func (c Clause) Valid(f func(VariableName, TermName) error) error {
	return f(c.Variable, c.Term)
}

func (c Clause) MarshalYAML() (interface{}, error) {
	return map[interface{}]interface{}{
		c.Variable: c.Term,
	}, nil
}

func C(variable VariableName, term TermName) Clause {
	return Clause{variable, term}
}

func (c Clause) String() string {
	return fmt.Sprintf("%s=%s", c.Variable, c.Term)
}

func (c Clause) Evaluate(fis FIS) num.Num {
	v := fis.GetInput(c.Variable)
	return v.GetTermValue(c.Term)
}

type Connector []Antecedent

func (a Connector) Valid(f func(VariableName, TermName) error) error {
	for _, clause := range a {
		if err := clause.Valid(f); err != nil {
			return err
		}
	}
	return nil
}

type And Connector

func (a And) Valid(f func(VariableName, TermName) error) error {
	return Connector(a).Valid(f)
}

func (a And) MarshalYAML() (interface{}, error) {
	return struct {
		And []Antecedent `yaml:"and"`
	}{a}, nil
}

func (c Connector) string(symbol string) string {
	var res []string
	for _, e := range c {
		res = append(res, fmt.Sprintf("%s", e))
	}
	return fmt.Sprintf("(%s)", strings.Join(res, symbol))
}

func (a And) Evaluate(fis FIS) num.Num {
	if len(a) == 0 {
		return num.NaN
	}
	res := a[0].Evaluate(fis)
	for _, b := range a[1:] {
		res = fis.And(res, b.Evaluate(fis))
	}
	return res
}

func (a And) String() string {
	return (Connector)(a).string(" and ")
}

type Or Connector

func (a Or) Valid(f func(VariableName, TermName) error) error {
	return Connector(a).Valid(f)
}

func (a Or) MarshalYAML() (interface{}, error) {
	return struct {
		Or []Antecedent `yaml:"or"`
	}{a}, nil
}

func (a Or) Evaluate(fis FIS) num.Num {
	if len(a) == 0 {
		return num.ZERO
	}
	res := a[0].Evaluate(fis)
	for _, b := range a[1:] {
		res = fis.Or(res, b.Evaluate(fis))
	}
	return res
}

func (a Or) String() string {
	return (Connector)(a).string(" or ")
}

type Not struct {
	Antecedent
}

var one = num.NewI(1, 0)

func (a Not) Evaluate(fis FIS) num.Num {
	return one.Sub(a.Antecedent.Evaluate(fis))
}

func (a Not) String() string {
	return fmt.Sprintf("not %s", a.Antecedent)
}
