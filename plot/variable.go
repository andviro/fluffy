package plot

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart/v2"

	"github.com/andviro/fluffy/v2"
	"github.com/andviro/fluffy/v2/num"
)

func Terms(fn string, xmin, xmax float64, terms []fluffy.Term) error {
	graph := chart.Chart{}
	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer f.Close()
	for _, v := range terms {
		s := chart.ContinuousSeries{
			Name: string(v.Name),
		}
		for x := xmin; x <= xmax; x += (xmax - xmin) / 100.0 {
			s.XValues = append(s.XValues, x)
			s.YValues = append(s.YValues, v.MembershipValue(num.NewF(x)).Float())
		}
		graph.Series = append(graph.Series, s)
	}
	graph.Elements = append(graph.Elements, chart.Legend(&graph))
	return graph.Render(chart.PNG, f)
}

func MembershipFunctions(fn string, src *fluffy.Variable) error {
	graph := chart.Chart{}
	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer f.Close()
	xmin, xmax := src.XMin.Float(), src.XMax.Float()
	if xmin == xmax {
		xmin, xmax = -10.0, 10.0
	}
	for _, v := range src.Terms {
		s := chart.ContinuousSeries{
			Name: string(v.Name),
		}
		for x := xmin; x <= xmax; x += (xmax - xmin) / 100.0 {
			s.XValues = append(s.XValues, x)
			s.YValues = append(s.YValues, v.MembershipValue(num.NewF(x)).Float())
		}
		graph.Series = append(graph.Series, s)
	}
	graph.Elements = append(graph.Elements, chart.Legend(&graph))
	return graph.Render(chart.PNG, f)
}
