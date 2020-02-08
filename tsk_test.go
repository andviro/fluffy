package fluffy

import (
	"testing"

	"github.com/andviro/fluffy/mf"
)

func TestTipper(t *testing.T) {
	tipper := TSK{
		Inputs: []Variable{
			{
				Name: "food",
				Terms: []Term{
					{
						Name:           "delicious",
						MembershipFunc: mf.RightLinear{5.5, 10.0},
					},
					{
						Name:           "rancid",
						MembershipFunc: mf.Triangle{0.0, 2.0, 5.5},
					},
				},
			},
			{
				Name: "service",
				Terms: []Term{
					{
						Name:           "excellent",
						MembershipFunc: mf.RightGaussian{10.0, 2.0},
					},
					{
						Name:           "good",
						MembershipFunc: mf.Gaussian{5.0, 2.0},
					},
					{
						Name:           "poor",
						MembershipFunc: mf.LeftGaussian{0.0, 2.0},
					},
				},
			},
		},
		Outputs: []TSKOutput{
			{
				Name: "tip",
				Terms: []TSKTerm{
					{
						Name:   "cheap",
						Coeffs: []float64{1.9, 5.6, 6.0},
					},
					{
						Name:   "generous",
						Coeffs: []float64{0.6, 1.3, 1.0},
					},
					{
						Name:   "average",
						Coeffs: []float64{1.6},
					},
				},
			},
		},
		Rules: []Rule{
			{
				Weight:      1.0,
				Antecedent:  Or{Clause{"food", "rancid"}, Clause{"service", "poor"}},
				Consequents: []Clause{{"tip", "cheap"}},
			},
			{
				Weight:      1.0,
				Antecedent:  Clause{"service", "good"},
				Consequents: []Clause{{"tip", "average"}},
			},
			{
				Weight:      1.0,
				Antecedent:  Or{Clause{"food", "delicious"}, Clause{"service", "excellent"}},
				Consequents: []Clause{{"tip", "generous"}},
			},
		},
	}
	tipper.SetInput("food", 8)
	tipper.SetInput("service", 3)
	tipper.Evaluate()
	t.Log(tipper.GetOutput("tip"))
}
