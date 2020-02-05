package fluffy

type CoreFIS interface {
	And(a, b float64) float64
	Or(a, b float64) float64
	Not(b float64) float64
	Imply(a, b float64) float64
}
