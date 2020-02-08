package fluffy

type Clause struct {
	Variable string
	Term     string
}

func (c Clause) Evaluate(fis FIS) float64 {
	v := fis.GetInput(c.Variable)
	return v.GetTermValue(c.Term)
}

type Connector struct {
	A, B Evaluator
}

type And Connector

func (a And) Evaluate(fis FIS) float64 {
	return fis.And(a.A.Evaluate(fis), a.B.Evaluate(fis))
}

type Or Connector

func (a Or) Evaluate(fis FIS) float64 {
	return fis.Or(a.A.Evaluate(fis), a.B.Evaluate(fis))
}

type Not struct {
	Evaluator
}

func (a Not) Evaluate(fis FIS) float64 {
	return 1 - a.Evaluator.Evaluate(fis)
}
