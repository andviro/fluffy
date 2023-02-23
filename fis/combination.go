package fis

import (
	"fmt"

	"github.com/andviro/fluffy"
)

type ExternalInput struct {
	FIS    fluffy.FIS
	Prefix fluffy.VariableName
	Output fluffy.VariableName
}

func (tsk *TSK) ConnectFISOutput(prefix, output fluffy.VariableName, fis fluffy.FIS, terms ...fluffy.Term) error {
	fullName := prefix + "_" + output
	for _, v := range tsk.Inputs {
		if fullName == v.Name {
			return fmt.Errorf("variable %s already exists in inputs", output)
		}
	}
	tsk.ExternalInputs = append(tsk.ExternalInputs, ExternalInput{FIS: fis, Output: output, Prefix: prefix})
	tsk.Inputs = append(tsk.Inputs, &fluffy.Variable{
		Name:  fullName,
		Terms: terms,
	})
	return nil
}
