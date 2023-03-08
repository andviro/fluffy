package fis_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/andviro/goldie"
	"gopkg.in/yaml.v2"

	"github.com/andviro/fluffy/v2"
	"github.com/andviro/fluffy/v2/fis"
	"github.com/andviro/fluffy/v2/mf"
	"github.com/andviro/fluffy/v2/num"
	"github.com/andviro/fluffy/v2/op"
	"github.com/andviro/fluffy/v2/plot"
)

var one = num.NewF(1)

var tipper = fis.TSK{
	OrMethod: op.Probor,
	Inputs: []*fluffy.Variable{
		{
			Name: "food",
			XMin: num.ZERO,
			XMax: num.NewF(10),
			Terms: []fluffy.Term{
				{
					Name:           "delicious",
					MembershipFunc: mf.RightLinear{A: num.NewF(7), B: num.NewF(9)},
				},
				{
					Name:           "rancid",
					MembershipFunc: mf.LeftLinear{A: one, B: num.NewF(3)},
				},
			},
		},
		{
			Name: "service",
			XMin: num.ZERO,
			XMax: num.NewF(10),
			Terms: []fluffy.Term{
				{
					Name:           "excellent",
					MembershipFunc: mf.RightGaussian{C: num.NewF(10.0), Sigma: num.NewF(1.5)},
				},
				{
					Name:           "good",
					MembershipFunc: mf.Gaussian{C: num.NewF(5.0), Sigma: num.NewF(1.5)},
				},
				{
					Name:           "poor",
					MembershipFunc: mf.LeftGaussian{C: num.ZERO, Sigma: num.NewF(1.5)},
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
					Coeffs: []num.Num{num.NewF(15)},
				},
				{
					Name:   "cheap",
					Coeffs: []num.Num{num.NewF(5)},
				},
				{
					Name:   "generous",
					Coeffs: []num.Num{num.NewF(25)},
				},
			},
		},
	},
	Rules: []fluffy.Rule{
		{
			Weight: one,
			Antecedent: fluffy.Or{
				fluffy.C("food", "rancid"),
				fluffy.C("service", "poor"),
			},
			Consequents: []fluffy.Clause{
				fluffy.C("tip", "cheap"),
			},
		},
		{
			Weight:     one,
			Antecedent: fluffy.C("service", "good"),
			Consequents: []fluffy.Clause{
				fluffy.C("tip", "average"),
			},
		},
		{
			Weight: one,
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

func TestTSK_Tipper(t *testing.T) {
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
	for _, r := range tipper.Rules {
		fmt.Fprintf(buf, "%s\n", r)
	}
	for _, tc := range []testCase{{1, 2}, {3, 5}, {2, 7}, {3, 1}, {1, 3}, {8, 3}, {3, 8}} {
		tipper.SetInput("service", num.NewF(tc.Service))
		tipper.SetInput("food", num.NewF(tc.Food))
		tipper.Evaluate()
		fmt.Fprintf(buf, "%v => %v\n", tc, tipper.GetOutput("tip"))
	}
	goldie.Assert(t, "tsk-tipper", buf.Bytes())
}

func TestTSK_MarshalYAML(t *testing.T) {
	data, err := yaml.Marshal(tipper)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", data)
}
