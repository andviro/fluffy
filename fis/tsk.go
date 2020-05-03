package fis

import (
	"fmt"
	"math"

	"github.com/andviro/fluffy"
	"github.com/andviro/fluffy/op"
)

type TSK struct {
	AndMethod op.Binary `yaml:"andMethod,omitempty"`
	OrMethod  op.Binary `yaml:"orMethod,omitempty"`
	Inputs    []*fluffy.Variable
	Outputs   []TSKOutput
	Rules     []fluffy.Rule
}

type TSKOutput struct {
	Name         fluffy.VariableName `yaml:"name"`
	Terms        []TSKTerm           `yaml:"terms"`
	DefaultValue float64             `yaml:"defaultValue"`
	evaluations  []wz
}

type wz struct {
	w, z float64
}

type TSKTerm struct {
	Name   fluffy.TermName `yaml:"name"`
	Coeffs []float64
	z      float64
}

func (t *TSKTerm) Evaluate(fis *TSK) {
	res := t.Coeffs[0]
	for i, k := range t.Coeffs[1:] {
		res += k * fis.Inputs[i].GetValue()
	}
	t.z = res
}

func (v TSKOutput) GetValue() float64 {
	if len(v.evaluations) == 0 {
		return v.DefaultValue
	}
	num, denom := 0.0, 0.0
	for _, wz := range v.evaluations {
		denom += wz.w
		num += wz.w * wz.z
	}
	return num / denom
}

func (v *TSKOutput) reset(fis *TSK) {
	v.evaluations = nil
	for j := range v.Terms {
		v.Terms[j].Evaluate(fis)
	}
}

var _ fluffy.FIS = (*TSK)(nil)

func (fis *TSK) And(a float64, b float64) float64 {
	if fis.AndMethod != nil {
		return fis.AndMethod(a, b)
	}
	return op.Min(a, b)
}

func (fis *TSK) Or(a float64, b float64) float64 {
	if fis.OrMethod != nil {
		return fis.OrMethod(a, b)
	}
	return op.Max(a, b)
}

func (fis *TSK) GetInput(name fluffy.VariableName) *fluffy.Variable {
	for _, i := range fis.Inputs {
		if i.Name == name {
			return i
		}
	}
	return &fluffy.Variable{Name: name}
}

func (fis *TSK) Activate(c fluffy.Clause, w float64) {
	for i, o := range fis.Outputs {
		if o.Name == c.Variable {
			for _, t := range o.Terms {
				if t.Name == c.Term {
					fis.Outputs[i].evaluations = append(fis.Outputs[i].evaluations, wz{w: w, z: t.z})
				}
			}
		}
	}
}

func (fis *TSK) Validate() error {
	for _, o := range fis.Outputs {
		for _, t := range o.Terms {
			if n := len(t.Coeffs); n != 1 && n != len(fis.Inputs)+1 {
				return fmt.Errorf("term %s of output %s has invalid number of coefficients (%d)", t.Name, o.Name, n)
			}
		}
	}
	return nil
}

func (fis *TSK) Evaluate() {
	for i := range fis.Outputs {
		fis.Outputs[i].reset(fis)
	}
	for _, r := range fis.Rules {
		r.Evaluate(fis)
	}
}

func (fis *TSK) SetInput(name fluffy.VariableName, value float64) {
	for i := range fis.Inputs {
		if fis.Inputs[i].Name == name {
			fis.Inputs[i].SetValue(value)
			break
		}
	}
}

func (fis *TSK) GetOutput(name fluffy.VariableName) float64 {
	for _, i := range fis.Outputs {
		if i.Name == name {
			return i.GetValue()
		}
	}
	return math.NaN()
}
