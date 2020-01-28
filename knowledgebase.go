package fluffy

type KnowledgeBase struct {
	Inputs  map[string]Variable
	Outputs map[string]Variable
	Rules   map[string]Rule
}

func (kb *KnowledgeBase) GetInput(name string) Variable {
	return kb.Inputs[name]
}
