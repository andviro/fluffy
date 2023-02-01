package fis

import (
	"fmt"

	"github.com/andviro/fluffy"
	"github.com/andviro/fluffy/num"
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
	DefaultValue num.Num             `yaml:"defaultValue"`
	evaluations  []wz
}

type wz struct {
	w, z num.Num
}

type TSKTerm struct {
	Name   fluffy.TermName `yaml:"name"`
	Coeffs []num.Num
	z      num.Num
}

func (t *TSKTerm) Evaluate(fis *TSK) {
	res := t.Coeffs[0]
	for i, k := range t.Coeffs[1:] {
		res = res.Add(k.Mul(fis.Inputs[i].GetValue()))
	}
	t.z = res
}

func (v TSKOutput) GetValue() num.Num {
	if len(v.evaluations) == 0 {
		return v.DefaultValue
	}
	num, denom := num.ZERO, num.ZERO
	for _, wz := range v.evaluations {
		denom = denom.Add(wz.w)
		num = num.Add(wz.w.Mul(wz.z))
	}
	return num.Div(denom)
}

func (v *TSKOutput) reset(fis *TSK) {
	v.evaluations = nil
	for j := range v.Terms {
		v.Terms[j].Evaluate(fis)
	}
}

var _ fluffy.FIS = (*TSK)(nil)

func (fis *TSK) And(a num.Num, b num.Num) num.Num {
	if fis.AndMethod != nil {
		return fis.AndMethod(a, b)
	}
	return op.Min(a, b)
}

func (fis *TSK) Or(a num.Num, b num.Num) num.Num {
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

func (fis *TSK) Activate(c fluffy.Clause, w num.Num) {
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
	for _, r := range fis.Rules {
		if err := r.Antecedent.Valid(func(name fluffy.VariableName) error {
			for _, i := range fis.Inputs {
				if i.Name == name {
					return nil
				}
			}
			return fmt.Errorf("variable %s not found in inputs", name)
		}); err != nil {
			return err
		}
		for _, c := range r.Consequents {
			if err := c.Valid(func(name fluffy.VariableName) error {
				for _, i := range fis.Outputs {
					if i.Name == name {
						return nil
					}
				}
				return fmt.Errorf("consequent %s not found in outputs", name)
			}); err != nil {
				return err
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

func (fis *TSK) SetInput(name fluffy.VariableName, value num.Num) {
	for i := range fis.Inputs {
		if fis.Inputs[i].Name == name {
			fis.Inputs[i].SetValue(value)
			break
		}
	}
}

func (fis *TSK) GetOutput(name fluffy.VariableName) num.Num {
	for _, i := range fis.Outputs {
		if i.Name == name {
			return i.GetValue()
		}
	}
	return num.NaN
}
