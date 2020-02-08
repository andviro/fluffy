package fluffy

import (
	"math"
)

type KnowledgeBase interface {
	And(a float64, b float64) float64
	Or(a float64, b float64) float64
	GetInput(name string) Variable
	Activate(c Clause, w float64)
}

type TSK struct {
	AndMethod func(float64, float64) float64
	OrMethod  func(float64, float64) float64
	Inputs    []Variable
	Outputs   []TSKOutput
	Rules     []Rule
}

type TSKOutput struct {
	Name        string
	Terms       []TSKTerm
	evaluations []WZ
}

type WZ struct {
	W, Z float64
}

type TSKTerm struct {
	Name   string
	Coeffs []float64
}

func (t TSKTerm) Evaluate(kb *TSK) float64 {
	if len(t.Coeffs) == 1 {
		return t.Coeffs[0]
	}
	res := t.Coeffs[len(t.Coeffs)-1]
	for i := len(t.Coeffs) - 2; i >= 0; i-- {
		res += t.Coeffs[i] * kb.Inputs[i].GetValue()
	}
	return res
}

func (v TSKOutput) GetValue() float64 {
	num, denom := 0.0, 0.0
	for _, wz := range v.evaluations {
		denom += wz.W
		num += wz.W * wz.Z
	}
	if denom == 0 {
		return math.NaN()
	}
	return num / denom
}

var _ KnowledgeBase = (*TSK)(nil)

func (kb *TSK) And(a float64, b float64) float64 {
	if kb.AndMethod != nil {
		return kb.AndMethod(a, b)
	}
	return math.Min(a, b)
}

func (kb *TSK) Or(a float64, b float64) float64 {
	if kb.OrMethod != nil {
		return kb.OrMethod(a, b)
	}
	return math.Max(a, b)
}

func (kb *TSK) GetInput(name string) Variable {
	for _, i := range kb.Inputs {
		if i.Name == name {
			return i
		}
	}
	return Variable{Name: name}
}

func (kb *TSK) Activate(c Clause, w float64) {
	for i, o := range kb.Outputs {
		if o.Name == c.Variable {
			for _, t := range o.Terms {
				if t.Name == c.Term {
					kb.Outputs[i].evaluations = append(kb.Outputs[i].evaluations, WZ{W: w, Z: t.Evaluate(kb)})
				}
			}
		}
	}
}

func (kb *TSK) Evaluate() {
	for i := range kb.Outputs {
		kb.Outputs[i].evaluations = nil
	}
	for _, r := range kb.Rules {
		r.Evaluate(kb)
	}
}

func (kb *TSK) SetInput(name string, value float64) {
	for i := range kb.Inputs {
		if kb.Inputs[i].Name == name {
			kb.Inputs[i].SetValue(value)
			break
		}
	}
}

func (kb *TSK) GetOutput(name string) float64 {
	for _, i := range kb.Outputs {
		if i.Name == name {
			return i.GetValue()
		}
	}
	return math.NaN()
}
