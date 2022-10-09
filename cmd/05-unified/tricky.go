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

// Parameters for the tricky graph test space.
const (
	aMin  = 20
	aMax  = 100
	aStep = 10

	kMin  = float32(0.1)
	kMax  = 1
	kStep = 0.1
)

func runTricky(
	comparison cover.Comparison,
	weigher graph.Weigher,
	nameStub string,
) *charts.HeatMap {
	name := fmt.Sprintf("%s-tricky", nameStub)
	pb := progressbar.New((aMax - aMin) / aStep)

	as := []int{}
	ks := []string{}
	series := []opts.HeatMapData{}

	aIndex := 0
	for a := aMin; a <= aMax; a += aStep {
		as = append(as, a)
		// BODGE: only want one set of ks.
		ks = []string{}

		for k := kMin; k <= kMax; k += kStep {
			kActual := int(math.Round(float64(k) * float64(a)))
			ks = append(ks, fmt.Sprintf("%.2f", k))
			pb.Describe(fmt.Sprintf("a=%d, k=%.2f, kn=%.2f", a, k, k*float32(a)))
			var sumDeltas float32

			for i := 0; i <= reps; i++ {
				g := graph.NewTricky(a, kActual, weigher)
				res := comparison.Run(g)
				sumDeltas += (res.Delta() / res.Baseline) * 100
			}

			mean := sumDeltas / reps
			series = append(series, opts.HeatMapData{
				Value: [3]interface{}{fmt.Sprintf("%.2f", k), aIndex, mean},
				Name:  fmt.Sprintf("a=%v, k=%v", a, kActual),
			})
		}

		pb.Add(1)
		aIndex++
	}

	heatMap := heatMapBase(name, as)
	heatMap.SetXAxis(ks).AddSeries("means", series)
	heatMap.Validate()

	page := components.NewPage()
	page.AddCharts(heatMap)

	f, _ := os.Create(fmt.Sprintf("out/%s.html", name))
	page.Render(io.MultiWriter(f))

	return heatMap
}
