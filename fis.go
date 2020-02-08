package fluffy

type FIS interface {
	And(a float64, b float64) float64
	Or(a float64, b float64) float64
	GetInput(name string) Variable
	Activate(c Clause, w float64)
}
