package main

import (
	"fmt"

	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

func main() {
	for p := float32(0); p <= 1; p += 0.1 {
		for n := 5; n <= 100; n += 5 {
			fmt.Printf("\nTesting n=%v, p=%v\n", n, p)

			// TODO: try a number of iterations.
			g := graph.NewWeighted(n, p, graph.Uniform(1))
			test(g)
		}
	}
}

func test(g *graph.Weighted) {
	// g.Print()
	strats := map[string]cover.Strategy{
		// "exhaustive": cover.NewExhaustive(g),
		"clever":   cover.NewClever(g),
		"vazirani": cover.NewVazirani(g),
	}
	outcomes := map[string]float32{}
	for name, strat := range strats {
		outcomes[name] = strat.CoverWeight()
	}
	fmt.Printf("OUTCOMES: %v\n", outcomes)
}
