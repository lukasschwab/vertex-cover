package main

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/schollz/progressbar/v3"
)

const (
	aMin  = 20
	aMax  = 100
	aStep = 10

	kMin  = float32(0.1)
	kMax  = 1
	kStep = 0.1

	// // Number of vertices.
	// nMin  = 10
	// nMax  = 100
	// nStep = 10

	// // Edge probability.
	// pMin  = float32(0.1)
	// pMax  = 1
	// pStep = 0.05
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
	run(comp, "Clever")

	comp.Test = func(g *graph.Weighted) cover.Strategy {
		return cover.NewLavrov(g)
	}
	run(comp, "Lavrov")
}

func run(c cover.Comparison, testName string) {
	ns := []int{}
	ps := []string{}
	series := []opts.HeatMapData{}

	pb := progressbar.New((aMax - aMin) / aStep)

	nIndex := 0
	for a := aMin; a <= aMax; a += aStep {
		ns = append(ns, a)
		pIndex := 0
		// BODGE: only want one set of ps.
		ps = []string{}

		for k := kMin; k <= kMax; k += kStep {
			kActual := int(math.Round(float64(k) * float64(a)))
			ps = append(ps, fmt.Sprintf("%.2f", k))
			pb.Describe(fmt.Sprintf("%v: a=%d, k=%.2f, kn=%.2f", testName, a, k, k*float32(a)))
			var sumDeltas float32
			for i := 0; i <= reps; i++ {
				g := graph.NewTricky(a, kActual, graph.Random)

				// g := graph.NewWeighted(n, p, graph.Uniform{})
				res := c.Run(g)
				// sumDeltas += res.Delta()
				// NOTE: trying normalization.
				if a == 10 && k == 0.1 {
					fmt.Printf("CUR RESULTS: %+v\n", res)
				}

				sumDeltas += (res.Delta() / res.Baseline) * 100
			}
			mean := sumDeltas / reps
			series = append(series, opts.HeatMapData{
				Value: [3]interface{}{fmt.Sprintf("%.2f", k), nIndex, mean},
				Name:  fmt.Sprintf("a=%v, k=%v", a, kActual),
			})
			pIndex++
		}
		pb.Add(1)
		nIndex++
	}

	heatMap := heatMapBase(testName, ns)
	fmt.Printf("K ratios: %v\n", ps)
	heatMap.SetXAxis(ps).AddSeries("means", series)

	heatMap.Validate()

	page := components.NewPage()
	page.AddCharts(heatMap)

	f, _ := os.Create(fmt.Sprintf("out/%v-vazirani.html", testName))
	if err := page.Render(io.MultiWriter(f)); err != nil {
		fmt.Printf("Error rendering: %v", err)
	}
}

func heatMapBase(testName string, yAxisData interface{}) *charts.HeatMap {
	hm := charts.NewHeatMap()
	hm.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("%v - Vazirani on tricky graphs", testName),
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:      "k ratio",
			Type:      "category",
			SplitArea: &opts.SplitArea{Show: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "a",
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
			Min:        -100,
			Max:        100,
			Text:       []string{"% Worse", "% Better"},
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)
	return hm
}
