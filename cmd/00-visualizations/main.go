package main

import (
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

func main() {
	page := components.NewPage()
	page.AddCharts(
		graph.GraphA.Display("GraphA"),
		graph.GraphB.Display("GraphB"),
		graph.GraphC.Display("GraphC"),
		graph.GraphVaziraniFail.Display("GraphVaziraniFail"),
	)
	f, err := os.Create("out/fixtures.html")
	if err != nil {
		panic(err)
	}
	page.Render(f)

	page = components.NewPage()
	page.AddCharts(
		graph.NewTricky(20, 5, graph.Uniform).Display("Lavrov's example"),
	)
	f, err = os.Create("out/tricky.html")
	if err != nil {
		panic(err)
	}
	page.Render(f)

	page = components.NewPage()
	for p := float32(0); p <= 1.0; p += 0.1 {
		page.AddCharts(
			graph.NewWeighted(10, p, graph.Uniform).
				Display(fmt.Sprintf("p=%.1f", p)),
		)
	}
	f, err = os.Create("out/random.html")
	if err != nil {
		panic(err)
	}
	page.Render(f)
}
