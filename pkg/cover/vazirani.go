package cover

import (
	"fmt"
	"math/rand"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
)

type vazirani struct {
	weight float32
	g      *graph.Weighted
	t      map[graph.Vertex]float32
}

func NewVazirani(g *graph.Weighted) Strategy {
	t := make(map[graph.Vertex]float32, len(g.Vertices()))
	for _, v := range g.Vertices() {
		t[v] = g.Weight(v)
	}

	return &vazirani{
		weight: 0,
		g:      g,
		t:      t,
	}
}

func (v *vazirani) CoverWeight() float32 {
	return v.search()
}

func (v *vazirani) search() float32 {
	v.g.Print()

	// While Vreq is not a vertex cover...
	if isCovered(v.g) {
		return v.weight
	}

	// Pick an uncovered edge, say (u, v).
	edges := v.g.Edges()
	fmt.Printf("Ts: %v\n", v.t)
	edge := edges[rand.Intn(len(edges))]
	vertexU, vertexV := edge[0], edge[1]

	fmt.Printf("%v (%v) -> %v (%v)\n\n", vertexU, v.g.Weight(vertexU), vertexV, v.g.Weight(vertexV))

	// Let m = min (t(u), t(v)).
	m := v.t[vertexV]
	if v.t[vertexU] < v.t[vertexV] {
		m = v.t[vertexU]
	}

	// t(u) ← t(u) − m.
	v.t[vertexU] -= m
	// t(v) ← t(v) − m.
	v.t[vertexV] -= m

	// Include in Vreq all vertices having t(v) = 0.
	if v.t[vertexU] == float32(0) {
		v.weight += v.g.Weight(vertexU)
		v.g = v.g.Without(vertexU)
	}
	if v.t[vertexV] == float32(0) {
		v.weight += v.g.Weight(vertexV)
		v.g = v.g.Without(vertexV)
	}

	return v.search()
}
