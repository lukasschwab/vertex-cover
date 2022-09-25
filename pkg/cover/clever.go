package cover

import "github.com/lukasschwab/vertex-cover/pkg/graph"

type clever struct {
	weight float32
	g      *graph.Weighted
}

func NewClever(g *graph.Weighted) Strategy {
	return clever{
		weight: 0,
		g:      g,
	}
}

func (c clever) CoverWeight() float32 {
	return c.search()
}

func (c clever) search() float32 {
	// If you have a cover, halt.
	if isCovered(c.g) {
		return c.weight
	}

	// Take v with the highest degree.
	var taken graph.Vertex
	var takenDegree int
	var takenWeight float32
	for _, v := range c.g.Vertices() {
		if degree := c.g.Degree(v); degree > takenDegree {
			taken, takenDegree, takenWeight = v, degree, c.g.Weight(v)
		} else if degree == takenDegree && c.g.Weight(v) < takenWeight {
			// Tie-break by weight.
			taken, takenDegree, takenWeight = v, degree, c.g.Weight(v)
		}
	}

	cprime := clever{
		weight: c.weight + c.g.Weight(taken),
		g:      c.g.Without(taken),
	}
	return cprime.search()
}

func isCovered(g *graph.Weighted) bool {
	for _, v := range g.Vertices() {
		if g.Neighbors(v).Length() != 0 {
			return false
		}
	}
	return true
}
