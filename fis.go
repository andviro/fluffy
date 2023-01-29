package fluffy

import num "github.com/robaho/fixed"

type FIS interface {
	And(a num.Fixed, b num.Fixed) num.Fixed
	Or(a num.Fixed, b num.Fixed) num.Fixed
	GetInput(name VariableName) *Variable
	Activate(c Clause, w num.Fixed)
}
