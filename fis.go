package fluffy

import "github.com/shopspring/decimal"

type FIS interface {
	And(a decimal.Decimal, b decimal.Decimal) decimal.Decimal
	Or(a decimal.Decimal, b decimal.Decimal) decimal.Decimal
	GetInput(name VariableName) *Variable
	Activate(c Clause, w decimal.Decimal)
}
