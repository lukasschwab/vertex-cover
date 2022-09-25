package cover

import (
	"fmt"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

// Exhaustive implements Strategy.
type Exhaustive struct {
	*graph.Weighted
}

func (e Exhaustive) CoverWeight() float32 {
	return e.search([]graph.Vertex{}, e.Vertices()).float32
}

func (e Exhaustive) search(included, candidates []graph.Vertex) Weight {
	if len(candidates) == 0 {
		// This is the bottom; check if it's a cover, and bubble up the weight
		// if it ain't.
		return e.weight(included)
	}

	head, tail := candidates[0], candidates[1:]

	weightWithout := e.search(included, tail)
	weightWith := e.search(append(included, head), tail)

	return weightWithout.OrLesser(weightWith)
}

func (e Exhaustive) weight(included []graph.Vertex) Weight {
	fmt.Printf("Evaluating cover: %v\n", included)

	includedSet := graph.NewNeighbors(included)

	var totalWeight float32 = 0
	for _, v := range e.Vertices() {
		// Either v is in the set cover...
		if includedSet.Includes(v) {
			totalWeight += e.Weight(v)
			continue
		}
		// ...or every neighbor is in the set cover.
		for _, neighbor := range e.Neighbors(v).Vertices() {
			if !includedSet.Includes(neighbor) {
				return Weight{noCover: true}
			}
		}
	}

	fmt.Printf("It covers; got weight: %v\n", totalWeight)

	// It's a cover! Calculate the weight.
	return Weight{float32: totalWeight}
}
