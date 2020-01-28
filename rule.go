package fluffy

type InputGetter interface {
	GetInput(name string) Variable
}

type Evaluator interface {
	Evaluate(src) float64
}

type RuleBase interface {
	And(float64, float64) float64
	Or(float64, float64) float64
}

type Clause struct {
	Variable string
	Term     string
}
