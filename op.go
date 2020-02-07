package fluffy

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
