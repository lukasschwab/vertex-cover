package cover

import (
	"math"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/schollz/progressbar/v3"
)

// exhaustive implements Strategy.
type exhaustive struct {
	*graph.Weighted
	bar *progressbar.ProgressBar
}

// This is useless. It'l take 15000 hours to check a graph with 45 vertices on
// my machine.
func NewExhaustive(g *graph.Weighted) Strategy {
	pfloat := math.Pow(2, float64(len(g.Vertices())))
	bar := progressbar.Default(int64(math.Round(pfloat)))

	return exhaustive{g, bar}
}

func (e exhaustive) CoverWeight() float32 {
	// TODO: validate !(Weight).noCover
	return e.search([]graph.Vertex{}, e.Vertices()).float32
}

func (e exhaustive) search(included, candidates []graph.Vertex) Weight {
	if len(candidates) == 0 {
		// This is the bottom; check if it's a cover, and bubble up the weight
		// if it ain't.
		e.bar.Add(1)
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
