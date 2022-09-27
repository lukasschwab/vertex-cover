package main

import (
	"fmt"

	"github.com/lukasschwab/vertex-cover/pkg/cover"
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

// main demonstrates Lavrov outperforming clever for relatively high k, because
// k determines the proportional "trickiness" of a graph for clever (how far it
// will perform from the optimum; the unbounded factor Hk).
func main() {
	for a := 20; a <= 100; a++ {
		for k := 5; k <= a; k++ {
			g := graph.NewTricky(a, k, graph.Uniform(1))
			// opt := 2 * a
			clever := cover.NewClever(g).CoverWeight()
			lavrov := cover.NewLavrov(g).CoverWeight()
			fmt.Printf("%v,", lavrov-clever)
		}
		fmt.Print("\n")
	}
}
