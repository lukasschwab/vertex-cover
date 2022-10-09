package main

import (
	"fmt"
	"log"

	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

// reps for which to run a given experiment.
const reps = 10

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
	"uniform":        graph.Uniform,
	"random":         graph.Random,
	"degreeNegative": graph.DegreeNegative,
	"degreePositive": graph.DegreePositive,
}

func main() {
	// pb := progressbar.New(len(comparisons) * len(weighers))
	for strategyName, comparison := range comparisons {
		for weigherName, weigher := range weighers {
			nameStub := fmt.Sprintf("%s-%s", strategyName, weigherName)
			log.Default().Printf("Testing %v...", nameStub)
			// pb.Describe("main: " + nameStub)
			run(comparison, weigher, nameStub)
			// pb.Add(1)
		}
	}
}

func run(comparison cover.Comparison, weigher graph.Weigher, nameStub string) {
	runRandom(comparison, weigher, nameStub)
	runTricky(comparison, weigher, nameStub)
}
