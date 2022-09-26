package cover

import (
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

// exhaustive implements Strategy.
type exhaustive struct {
	*graph.Weighted
}

func NewExhaustive(g *graph.Weighted) Strategy {
	return exhaustive{g}
}

func (e exhaustive) CoverWeight() float32 {
	// TODO: validate !(Weight).noCover
	return e.search([]graph.Vertex{}, e.Vertices()).float32
}

func (e exhaustive) search(included, candidates []graph.Vertex) Weight {
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

func (e exhaustive) weight(included []graph.Vertex) Weight {
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

	// It's a cover! Calculate the weight.
	return Weight{float32: totalWeight}
}
