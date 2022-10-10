package main

import (
	"fmt"
	"log"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

// reps for which to run a given experiment.
const reps = 10

var logger = log.Default()

// baselineStrategy for all comparisons: always Vazirani.
func baselineStrategy(g *graph.Weighted) cover.Strategy {
	return cover.NewVazirani(g)
}

var comparisons = map[string]cover.Comparison{
	"clever": {
		Baseline: baselineStrategy,
		Test: func(g *graph.Weighted) cover.Strategy {
			return cover.NewClever(g)
		},
	},
	"lavrov": {
		Baseline: baselineStrategy,
		Test: func(g *graph.Weighted) cover.Strategy {
			return cover.NewLavrov(g)
		},
	},
}

var weighers = map[string]graph.Weigher{
	"uniform":                   graph.Uniform,
	"random":                    graph.Random,
	"degreeNegative":            graph.DegreeNegative,
	"degreePositive":            graph.DegreePositive,
	"degreePositiveSuperlinear": graph.DegreePositiveSuperlinear,
}

func main() {
	allHeatMaps := []components.Charter{}

	for strategyName, comparison := range comparisons {
		for weigherName, weigher := range weighers {
			nameStub := fmt.Sprintf("%s-%s", strategyName, weigherName)
			logger.Printf("Testing %v...", nameStub)
			random, tricky := run(comparison, weigher, nameStub)
			allHeatMaps = append(allHeatMaps, random, tricky)
		}
	}
	write("unified", allHeatMaps...)
}

func run(comparison cover.Comparison, weigher graph.Weigher, nameStub string) (random, tricky *charts.HeatMap) {
	return runRandom(comparison, weigher, nameStub), runTricky(comparison, weigher, nameStub)
}
