package fluffy

type knowledgeBase struct {
	CoreFIS
	Inputs  map[string]Variable
	Outputs map[string]Variable
	Rules   map[float64]Rule
}

func (kb *knowledgeBase) GetInput(name string) Variable {
	return kb.Inputs[name]
}

func (kb *knowledgeBase) SetInput(name string, value float64) {
	kb.Inputs[name] = kb.Inputs[name].SetValue(value)
}

func (kb *knowledgeBase) evaluate() map[string]map[string][]float64 {
	res := make(map[string]map[string][]float64)
	for strength, rule := range kb.Rules {
		val := strength * rule.Antecedent.Evaluate(kb)
		for _, c := range rule.Consequents {
			mf, ok := res[c.Variable]
			if !ok {
				mf = make(map[string][]float64)
			}
			mf[c.Term] = append(mf[c.Term], val)
			res[c.Variable] = mf
		}
	}
	return res
}

type sugeno struct {
	knowledgeBase
}
