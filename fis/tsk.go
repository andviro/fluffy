package fis

import (
	"fmt"

	"github.com/andviro/fluffy/v2"
	"github.com/andviro/fluffy/v2/num"
	"github.com/andviro/fluffy/v2/op"
)

type TSK[T num.Num[T]] struct {
	AndMethod      op.Binary[T] `yaml:"andMethod,omitempty"`
	OrMethod       op.Binary[T] `yaml:"orMethod,omitempty"`
	Inputs         []*fluffy.Variable[T]
	ExternalInputs []ExternalInput[T]
	Outputs        []TSKOutput[T]
	Rules          []fluffy.Rule[T]
}

type TSKOutput[T num.Num[T]] struct {
	Name         fluffy.VariableName `yaml:"name"`
	Terms        []TSKTerm[T]        `yaml:"terms"`
	DefaultValue T                   `yaml:"defaultValue"`
	evaluations  []wz[T]
}

type wz[T num.Num[T]] struct {
	w, z T
}

type TSKTerm[T num.Num[T]] struct {
	Name   fluffy.TermName `yaml:"name"`
	Coeffs []T
	z      T
}

func (t *TSKTerm[T]) Evaluate(fis *TSK[T]) {
	res := t.Coeffs[0]
	for i, k := range t.Coeffs[1:] {
		res = res.Add(k.Mul(fis.Inputs[i].GetValue()))
	}
	t.z = res
}

func (v TSKOutput[T]) GetValue() T {
	if len(v.evaluations) == 0 {
		return v.DefaultValue
	}
	var num, denom T
	for _, wz := range v.evaluations {
		denom = denom.Add(wz.w)
		num = num.Add(wz.w.Mul(wz.z))
	}
	return num.Div(denom)
}

func (v *TSKOutput[T]) reset(fis *TSK[T]) {
	v.evaluations = nil
	for j := range v.Terms {
		v.Terms[j].Evaluate(fis)
	}
}

func (fis *TSK[T]) And(a T, b T) T {
	if fis.AndMethod != nil {
		return fis.AndMethod(a, b)
	}
	return op.Min(a, b)
}

func (fis *TSK[T]) Or(a T, b T) T {
	if fis.OrMethod != nil {
		return fis.OrMethod(a, b)
	}
	return op.Max(a, b)
}

func (fis *TSK[T]) GetInput(name fluffy.VariableName) *fluffy.Variable[T] {
	for _, i := range fis.Inputs {
		if i.Name == name {
			return i
		}
	}
	return &fluffy.Variable[T]{Name: name}
}

func (fis *TSK[T]) Activate(c fluffy.Clause[T], w T) {
	for i, o := range fis.Outputs {
		if o.Name == c.Variable {
			for _, t := range o.Terms {
				if t.Name == c.Term {
					fis.Outputs[i].evaluations = append(fis.Outputs[i].evaluations, wz[T]{w: w, z: t.z})
				}
			}
		}
	}
}

func (fis *TSK[T]) Validate() error {
	for _, o := range fis.Outputs {
		for _, t := range o.Terms {
			if n := len(t.Coeffs); n != 1 && n != len(fis.Inputs)+1 {
				return fmt.Errorf("term %s of output %s has invalid number of coefficients (%d)", t.Name, o.Name, n)
			}
		}
	}
	for _, r := range fis.Rules {
		if err := r.Antecedent.Valid(func(variable fluffy.VariableName, term fluffy.TermName) error {
			for _, i := range fis.Inputs {
				if i.Name == variable {
					for _, t := range i.Terms {
						if t.Name == term {
							return nil
						}
					}
					return fmt.Errorf("term value %s not found for input %s", term, variable)
				}
			}
			return fmt.Errorf("variable %s not found in inputs", variable)
		}); err != nil {
			return err
		}
		for _, c := range r.Consequents {
			if err := c.Valid(func(name fluffy.VariableName, term fluffy.TermName) error {
				for _, o := range fis.Outputs {
					if o.Name == name {
						for _, t := range o.Terms {
							if t.Name == term {
								return nil
							}
						}
						return fmt.Errorf("term %s not found in consequent %s terms", term, name)
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

func (fis *TSK[T]) Evaluate() {
	for i := range fis.Outputs {
		fis.Outputs[i].reset(fis)
	}
	for _, ei := range fis.ExternalInputs {
		fullName := ei.Prefix + "_" + ei.Output
		fis.SetInput(fullName, ei.FIS.GetOutput(ei.Output))
	}
	for _, r := range fis.Rules {
		r.Evaluate(fis)
	}
}

func (fis *TSK[T]) SetInput(name fluffy.VariableName, value T) {
	for i := range fis.Inputs {
		if fis.Inputs[i].Name == name {
			fis.Inputs[i].SetValue(value)
			break
		}
	}
}

func (fis *TSK[T]) GetOutput(name fluffy.VariableName) T {
	for _, i := range fis.Outputs {
		if i.Name == name {
			return i.GetValue()
		}
	}
	return num.NaN[T]()
}
