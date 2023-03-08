package fluffy

import num "github.com/andviro/fluffy/v2/num"

type FIS interface {
	And(a num.Num, b num.Num) num.Num
	Or(a num.Num, b num.Num) num.Num
	GetInput(name VariableName) *Variable
	Activate(c Clause, w num.Num)
	GetOutput(name VariableName) num.Num
}
