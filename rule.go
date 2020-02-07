package fluffy

type Evaluator interface {
	Evaluate(kb KnowledgeBase) float64
}

type Rule struct {
	Weight      float64
	AndMethod   func(float64, float64) float64
	OrMethod    func(float64, float64) float64
	Antecedent  Evaluator
	Consequents []Clause
	KnowledgeBase
}

func (r *Rule) And(a float64, b float64) float64 {
	if r.AndMethod != nil {
		return r.AndMethod(a, b)
	}
	return r.KnowledgeBase.And(a, b)
}

func (r *Rule) Or(a float64, b float64) float64 {
	if r.OrMethod != nil {
		return r.OrMethod(a, b)
	}
	return r.KnowledgeBase.Or(a, b)
}

func (r *Rule) Evaluate(kb KnowledgeBase) {
	r.KnowledgeBase = kb
	w := r.Antecedent.Evaluate(r)
	for _, c := range r.Consequents {
		kb.Activate(c, w*r.Weight)
	}
}
