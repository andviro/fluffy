package fluffy

type Clause struct {
	Variable string
	Term     string
}

func (c Clause) Evaluate(kb KnowledgeBase) float64 {
	v := kb.GetInput(c.Variable)
	return v.GetTermValue(c.Term)
}
