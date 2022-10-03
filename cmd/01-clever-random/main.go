package main

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/schollz/progressbar/v3"
)

const (
	// Number of vertices.
	nMin  = 10
	nMax  = 100
	nStep = 10

	// Edge probability.
	pMin  = float32(0.1)
	pMax  = 1
	pStep = 0.05
)

// Number of random graphs to run against
const reps = 10

func main() {
	comp := cover.Comparison{
		Baseline: func(g *graph.Weighted) cover.Strategy {
			return cover.NewVazirani(g)
		},
		Test: func(g *graph.Weighted) cover.Strategy {
			return cover.NewClever(g)
		},
	}

	ns := []int{}
	ps := []string{}
	series := []opts.HeatMapData{}

	pb := progressbar.New((nMax - nMin) / nStep)
	nIndex := 0
	for n := nMin; n <= nMax; n += nStep {
		ns = append(ns, n)
		pIndex := 0
		// BODGE: only want one set of ps.
		ps = []string{}
		for p := pMin; p <= pMax; p += pStep {
			ps = append(ps, fmt.Sprintf("%.2f", p))
			pb.Describe(fmt.Sprintf("n=%d, p=%.2f", n, p))
			var sumDeltas float32 = 0
			for i := 0; i <= reps; i++ {
				g := graph.NewWeighted(n, p, graph.Uniform{})
				res := comp.Run(g)
				// sumDeltas += res.Delta()
				// NOTE: trying normalization.
				sumDeltas += (res.Delta() / res.Baseline) * 100
			}
			mean := sumDeltas / reps
			series = append(series, opts.HeatMapData{Value: [3]interface{}{fmt.Sprintf("%.2f", p), nIndex, mean}})
			pIndex++
		}
		pb.Add(1)
		nIndex++
	}

	heatMap := heatMapBase(ns)
	fmt.Printf("PS: %v\n", ps)
	heatMap.SetXAxis(ps).AddSeries("means", series)

	heatMap.Validate()

	page := components.NewPage()
	page.AddCharts(heatMap)

	f, _ := os.Create("out/clever-vazirani.html")
	page.Render(io.MultiWriter(f))
}

func heatMapBase(yAxisData interface{}) *charts.HeatMap {
	hm := charts.NewHeatMap()
	hm.SetGlobalOptions(
		// charts.WithTitleOpts(opts.Title{
		// 	Title: "Clever - Vazirani on random graphs",
		// }),
		charts.WithXAxisOpts(opts.XAxis{
			Name:      "p",
			Type:      "category",
			SplitArea: &opts.SplitArea{Show: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "n",
			Type: "category",
			Data: yAxisData,
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
			SplitArea: &opts.SplitArea{Show: true},
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Min:        -50,
			Max:        50,
			Text:       []string{"% Worse", "% Better"},
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)
	return hm
}
