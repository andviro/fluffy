package fluffy

import num "github.com/andviro/fluffy/v2/num"

type FIS[T num.Num[T]] interface {
	And(a T, b T) T
	Or(a T, b T) T
	GetInput(name VariableName) *Variable[T]
	Activate(c Clause[T], w T)
	GetOutput(name VariableName) T
}
