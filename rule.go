package fluffy

type KnowledgeBase interface {
	And(a float64, b float64) float64
	Or(a float64, b float64) float64
	GetInput(name string) Variable
}

type Evaluator interface {
	Evaluate(src KnowledgeBase) float64
}

type Clause struct {
	Variable string
	Term     string
}

func (c Clause) Evaluate(src KnowledgeBase) float64 {
	return src.GetInput(c.Variable).termValues[c.Term]
}

type Rule struct {
	Antecedent  Evaluator
	Consequents []Clause
}

type Binary struct {
	A, B Evaluator
}

type And Binary

func (a And) Evaluate(src KnowledgeBase) float64 {
	return src.And(a.A.Evaluate(src), a.B.Evaluate(src))
}

type Or Binary

func (a Or) Evaluate(src KnowledgeBase) float64 {
	return src.Or(a.A.Evaluate(src), a.B.Evaluate(src))
}
