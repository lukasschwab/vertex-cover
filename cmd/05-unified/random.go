package main

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/schollz/progressbar/v3"
)

// Parameters for the random graph test space.
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

func runRandom(
	comparison cover.Comparison,
	weigher graph.Weigher,
	nameStub string,
) *charts.HeatMap {
	name := fmt.Sprintf("%s-random", nameStub)
	pb := progressbar.New((nMax - nMin) / nStep)

	// Deposit experiment info for graphing.
	ns := []int{}
	ps := []string{}
	series := []opts.HeatMapData{}

	nIndex := 0
	for n := nMin; n <= nMax; n += nStep {
		ns = append(ns, n)
		pIndex := 0
		// BODGE: only want one set of ps. Don't need to re-record it every time.
		ps = []string{}

		for p := pMin; p <= pMax; p += pStep {
			ps = append(ps, fmt.Sprintf("%.2f", p))
			pb.Describe(fmt.Sprintf("n=%d, p=%.2f", n, p))
			var sumDeltas float32

			for i := 0; i <= reps; i++ {
				g := graph.NewWeighted(n, p, weigher)
				res := comparison.Run(g)
				sumDeltas += (res.Delta() / res.Baseline) * 100
			}

			mean := sumDeltas / reps
			series = append(series, opts.HeatMapData{Value: [3]interface{}{fmt.Sprintf("%.2f", p), nIndex, mean}})
			pIndex++
		}

		pb.Add(1)
		nIndex++
	}

	heatMap := heatMapBase(name, "p", "n", ns)
	heatMap.SetXAxis(ps).AddSeries("means", series)
	heatMap.Validate()
	write(name, heatMap)

	return heatMap
}
