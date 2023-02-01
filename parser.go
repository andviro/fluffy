package fluffy

import (
	"github.com/alecthomas/participle/v2"
	"github.com/andviro/fluffy/num"
)

type RulesParser struct {
	Rules []RuleParser `parser:"@@*"`
}

type RuleParser struct {
	Weight      string   `parser:"@Float"`
	Antecedent  OrParser `parser:"':' @@"`
	Consequents []Clause `parser:" ['-' '>'] @@+"`
}

func (rp RuleParser) Build() Rule {
	return Rule{
		Weight:      num.NewS(rp.Weight),
		Antecedent:  rp.Antecedent.Build(),
		Consequents: rp.Consequents,
	}
}

type OrParser struct {
	And  AndParser `parser:"@@"`
	Next *OrParser `parser:"[ '|' @@ ]"`
}

func (op OrParser) Build() Antecedent {
	res := []Antecedent{op.And.Build()}
	for next := op.Next; next != nil; next = next.Next {
		res = append(res, next.Build())
	}
	if len(res) == 1 {
		return res[0]
	}
	return Or(res)
}

func (ap AndParser) Build() Antecedent {
	res := []Antecedent{ap.Negation.Build()}
	for next := ap.Next; next != nil; next = next.Next {
		res = append(res, next.Build())
	}
	if len(res) == 1 {
		return res[0]
	}
	return And(res)
}

type AndParser struct {
	Negation NotParser  `parser:"@@"`
	Next     *AndParser `parser:"[ '&' @@ ]"`
}

func (np NotParser) Build() Antecedent {
	if np.Not != nil {
		return Not{np.Not.Build()}
	}
	return np.Next.Build()
}

type NotParser struct {
	Not  *NotParser    `parser:"( '!' @@ )"`
	Next *ClauseParser `parser:"| @@"`
}

func (cp ClauseParser) Build() Antecedent {
	if cp.Next != nil {
		return cp.Next.Build()
	}
	return Clause{VariableName(cp.Variable), TermName(cp.Term)}
}

type ClauseParser struct {
	Variable string    `parser:"@Ident '='"`
	Term     string    `parser:"@Ident"`
	Next     *OrParser `parser:"| '(' @@ ')'"`
}

var parser = participle.MustBuild[RulesParser](participle.UseLookahead(2))

func ParseRules(src string) ([]Rule, error) {
	expr, err := parser.ParseString("", src)
	if err != nil {
		return nil, err
	}
	var res []Rule
	for _, r := range expr.Rules {
		res = append(res, r.Build())
	}
	return res, nil
}
