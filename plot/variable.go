package plot

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart"

	"github.com/andviro/fluffy"
)

func MembershipFunctions(fn string, src fluffy.Variable) error {
	graph := chart.Chart{}
	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer f.Close()
	xmin, xmax := src.XMin, src.XMax
	for _, v := range src.Terms {
		s := chart.ContinuousSeries{
			Name: string(v.Name),
		}
		for x := xmin; x <= xmax; x += (xmax - xmin) / 100.0 {
			s.XValues = append(s.XValues, x)
			s.YValues = append(s.YValues, v.MembershipValue(x))
		}
		graph.Series = append(graph.Series, s)
	}
	graph.Elements = append(graph.Elements, chart.Legend(&graph))
	return graph.Render(chart.PNG, f)
}
