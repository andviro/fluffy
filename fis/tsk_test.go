package fis_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/andviro/goldie"

	"github.com/andviro/fluffy"
	"github.com/andviro/fluffy/fis"
	"github.com/andviro/fluffy/mf"
	"github.com/andviro/fluffy/op"
	"github.com/andviro/fluffy/plot"
)

func TestTSK_Tipper(t *testing.T) {
	tipper := fis.TSK{
		OrMethod: op.Probor,
		Inputs: []fluffy.Variable{
			{
				Name: "food",
				XMin: 0,
				XMax: 10,
				Terms: []fluffy.Term{
					{
						Name:           "delicious",
						MembershipFunc: mf.RightLinear{A: 7, B: 9},
					},
					{
						Name:           "rancid",
						MembershipFunc: mf.LeftLinear{A: 1, B: 3},
					},
				},
			},
			{
				Name: "service",
				XMin: 0,
				XMax: 10,
				Terms: []fluffy.Term{
					{
						Name:           "excellent",
						MembershipFunc: mf.RightGaussian{C: 10.0, Sigma: 1.5},
					},
					{
						Name:           "good",
						MembershipFunc: mf.Gaussian{C: 5.0, Sigma: 1.5},
					},
					{
						Name:           "poor",
						MembershipFunc: mf.LeftGaussian{C: 0.0, Sigma: 1.5},
					},
				},
			},
		},
		Outputs: []fis.TSKOutput{
			{
				Name: "tip",
				Terms: []fis.TSKTerm{
					{
						Name:   "average",
						Coeffs: []float64{15},
					},
					{
						Name:   "cheap",
						Coeffs: []float64{5},
					},
					{
						Name:   "generous",
						Coeffs: []float64{25},
					},
				},
			},
		},
		Rules: []fluffy.Rule{
			{
				Weight: 1.0,
				Antecedent: fluffy.Or{
					fluffy.C("food", "rancid"),
					fluffy.C("service", "poor"),
				},
				Consequents: []fluffy.Clause{
					fluffy.C("tip", "cheap"),
				},
			},
			{
				Weight:     1.0,
				Antecedent: fluffy.C("service", "good"),
				Consequents: []fluffy.Clause{
					fluffy.C("tip", "average"),
				},
			},
			{
				Weight: 1.0,
				Antecedent: fluffy.Or{
					fluffy.C("food", "delicious"),
					fluffy.C("service", "excellent"),
				},
				Consequents: []fluffy.Clause{
					fluffy.C("tip", "generous"),
				},
			},
		},
	}
	if err := tipper.Validate(); err != nil {
		t.Fatal(err)
	}
	for _, v := range tipper.Inputs {
		if err := plot.MembershipFunctions(filepath.Join("fixtures", fmt.Sprintf("%s.png", v.Name)), v); err != nil {
			t.Fatal(err)
		}
	}
	type testCase struct {
		Service float64
		Food    float64
	}
	buf := new(bytes.Buffer)
	for _, tc := range []testCase{{1, 2}, {3, 5}, {2, 7}, {3, 1}, {1, 3}, {8, 3}, {3, 8}} {
		tipper.SetInput("service", tc.Service)
		tipper.SetInput("food", tc.Food)
		tipper.Evaluate()
		fmt.Fprintf(buf, "%v => %f\n", tc, tipper.GetOutput("tip"))
	}
	goldie.Assert(t, "tsk-tipper", buf.Bytes())
}
