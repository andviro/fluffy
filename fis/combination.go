package fis

import (
	"fmt"

	fluffy "github.com/andviro/fluffy/v2"
	"github.com/andviro/fluffy/v2/num"
)

type ExternalInput[T num.Num[T]] struct {
	FIS    fluffy.FIS[T]
	Prefix fluffy.VariableName
	Output fluffy.VariableName
}

func (tsk *TSK[T]) ConnectFISOutput(prefix, output fluffy.VariableName, fis fluffy.FIS[T], terms ...fluffy.Term[T]) error {
	fullName := prefix + "_" + output
	for _, v := range tsk.Inputs {
		if fullName == v.Name {
			return fmt.Errorf("variable %s already exists in inputs", output)
		}
	}
	tsk.ExternalInputs = append(tsk.ExternalInputs, ExternalInput[T]{FIS: fis, Output: output, Prefix: prefix})
	tsk.Inputs = append(tsk.Inputs, &fluffy.Variable[T]{
		Name:  fullName,
		Terms: terms,
	})
	return nil
}
