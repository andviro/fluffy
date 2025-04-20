package fluffy

import (
	"github.com/alecthomas/participle/v2"
	"github.com/andviro/fluffy/v2/num"
)

type RulesParser[T num.Num[T]] struct {
	Rules []RuleParser[T] `parser:"@@*"`
}

type RuleParser[T num.Num[T]] struct {
	Weight      string      `parser:"(@Float | @Int)"`
	Antecedent  OrParser[T] `parser:"':' @@"`
	Consequents []Clause[T] `parser:" ['-' '>'] @@+"`
}

func (rp RuleParser[T]) Build() Rule[T] {
	return Rule[T]{
		Weight:      num.NewS[T](rp.Weight),
		Antecedent:  rp.Antecedent.Build(),
		Consequents: rp.Consequents,
	}
}

type OrParser[T num.Num[T]] struct {
	And  AndParser[T] `parser:"@@"`
	Next *OrParser[T] `parser:"[ '|' @@ ]"`
}

func (op OrParser[T]) Build() Antecedent[T] {
	res := []Antecedent[T]{op.And.Build()}
	for next := op.Next; next != nil; next = next.Next {
		res = append(res, next.Build())
	}
	if len(res) == 1 {
		return res[0]
	}
	return Or[T](res)
}

func (ap AndParser[T]) Build() Antecedent[T] {
	res := []Antecedent[T]{ap.Negation.Build()}
	for next := ap.Next; next != nil; next = next.Next {
		res = append(res, next.Build())
	}
	if len(res) == 1 {
		return res[0]
	}
	return And[T](res)
}

type AndParser[T num.Num[T]] struct {
	Negation NotParser[T]  `parser:"@@"`
	Next     *AndParser[T] `parser:"[ '&' @@ ]"`
}

func (np NotParser[T]) Build() Antecedent[T] {
	if np.Not != nil {
		return Not[T]{np.Not.Build()}
	}
	return np.Next.Build()
}

type NotParser[T num.Num[T]] struct {
	Not  *NotParser[T]    `parser:"( '!' @@ )"`
	Next *ClauseParser[T] `parser:"| @@"`
}

func (cp ClauseParser[T]) Build() Antecedent[T] {
	if cp.Next != nil {
		return cp.Next.Build()
	}
	return Clause[T]{VariableName(cp.Variable), TermName(cp.Term)}
}

type ClauseParser[T num.Num[T]] struct {
	Variable string       `parser:"@Ident '='"`
	Term     string       `parser:"@Ident"`
	Next     *OrParser[T] `parser:"| '(' @@ ')'"`
}

func parser[T num.Num[T]]() *participle.Parser[RulesParser[T]] {
	return participle.MustBuild[RulesParser[T]](participle.UseLookahead(2))
}

func ParseRules[T num.Num[T]](src string) ([]Rule[T], error) {
	expr, err := parser[T]().ParseString("", src)
	if err != nil {
		return nil, err
	}
	var res []Rule[T]
	for _, r := range expr.Rules {
		res = append(res, r.Build())
	}
	return res, nil
}
