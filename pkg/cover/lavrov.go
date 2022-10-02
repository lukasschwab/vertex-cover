package cover

import (
	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

type lavrov struct {
	weight float32
	g      *graph.Weighted
}

// NewLavrov returns a Strategy implementing Lavrov's aglorithm for unweighted
// vertex cover, described in
//
// + "Lecture 36: Approximation Algorithms" (2020)
// + "Graphs at Work:" http://lukasschwab.me/blog/gen/graphs-at-work.html
func NewLavrov(g *graph.Weighted) Strategy {
	return lavrov{
		weight: 0,
		g:      g,
	}
}

func (l lavrov) CoverWeight() float32 {
	return l.search()
}

func (l lavrov) search() float32 {
	if isCovered(l.g) {
		return l.weight
	}

	// Take the vertices on both sides of a random edge.
	var edge [2]graph.Vertex
	maxIncidentDegree := 0
	for _, e := range l.g.Edges() {
		degree := l.g.Degree(e[0]) + l.g.Degree(e[1])
		if degree > maxIncidentDegree {
			edge = e
			maxIncidentDegree = degree
		}
	}

	for _, v := range edge {
		l.weight += l.g.Weight(v)
		l.g = l.g.Without(v)
	}

	return l.search()
}
