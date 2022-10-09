package main

import (
	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

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
	"uniform":        graph.Uniform{},
	"random":         graph.Random{},
	"degreeNegative": graph.DegreeNegative{},
	"degreePositive": graph.DegreePositive{},
}

func main() {
	return
}
