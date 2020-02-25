package fluffy

import "math"

type Clause struct {
	Variable string
	Term     string
}

func (c Clause) Evaluate(fis FIS) float64 {
	v := fis.GetInput(c.Variable)
	return v.GetTermValue(c.Term)
}

type Connector []Evaluator

type And Connector

func (a And) Evaluate(fis FIS) float64 {
	if len(a) == 0 {
		return math.NaN()
	}
	res := a[0].Evaluate(fis)
	for _, b := range a[1:] {
		res = fis.And(res, b.Evaluate(fis))
	}
	return res
}

type Or Connector

func (a Or) Evaluate(fis FIS) float64 {
	if len(a) == 0 {
		return math.NaN()
	}
	res := a[0].Evaluate(fis)
	for _, b := range a[1:] {
		res = fis.Or(res, b.Evaluate(fis))
	}
	return res
}

type Not struct {
	Evaluator
}

func (a Not) Evaluate(fis FIS) float64 {
	return 1 - a.Evaluator.Evaluate(fis)
}
