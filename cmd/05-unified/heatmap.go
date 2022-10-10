package main

import (
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// heatMapBase returns a baseline heatmap for any test.
func heatMapBase(name string, yAxisData interface{}) *charts.HeatMap {
	hm := charts.NewHeatMap()
	hm.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: name,
		}),
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

// write a page of graphs.
func write(name string, graphs ...components.Charter) {
	logger.Printf("Rendering %v page...", name)
	page := components.NewPage().AddCharts(graphs...)
	if f, err := os.Create(fmt.Sprintf("out/%s.html", name)); err != nil {
		logger.Fatalf("Couldn't create output file '%s': %v", name, err)
	} else if err := page.Render(f); err != nil {
		logger.Printf("Error writing graph '%s': %v", name, err)
	}
}
