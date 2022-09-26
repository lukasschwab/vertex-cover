package main

import (
	"fmt"

	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

const (
	// reps for a given experiment (n, p)
	reps = 5
)

func main() {
	data := make([][]float32, 10)
	for i := range data {
		data[i] = make([]float32, 20)
	}

	i := 0
	for p := float32(0.1); p <= 1.1; p += 0.1 {
		j := 0
		for n := 5; n <= 100; n += 5 {
			for rep := 0; rep < reps; rep++ {
				// g := graph.NewTricky(20, 5, graph.Uniform(1))
				g := graph.NewWeighted(n, p, graph.Uniform(1))
				outcomes := test(g)
				data[i][j] += (outcomes["clever"] - outcomes["vazirani"])
			}
			j++
		}
		i++
	}

	// fmt.Println("%v\n", data)
	for _, row := range data {
		for _, sum := range row {
			fmt.Printf("%.3f,", sum/reps)
		}
		fmt.Printf("\n")
	}
}

func test(g *graph.Weighted) map[string]float32 {
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
	return outcomes
}
