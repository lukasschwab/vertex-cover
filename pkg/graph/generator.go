package graph

import "math/rand"

func NewUnweighted(n int, p float32) *Unweighted {
	g := &Unweighted{
		vertices: make(map[Vertex]*Neighbors, n),
	}
	// Initiate vertices, empty neighbors.
	for i := Vertex(0); i < Vertex(n); i++ {
		g.vertices[i] = NewNeighbors([]Vertex{})
	}

	for _, u := range g.Vertices() {
		for _, v := range g.Vertices() {
			if u <= v {
				if rand.Float32() < p {
					// TODO: refactor. Blech.
					g.vertices[u].vertices[v] = struct{}{}
					g.vertices[v].vertices[u] = struct{}{}
				}
			}
		}
	}

	return g
}

// type weigher func(Vertex, *Neighbors) float32

type weigher interface {
	weigh(Vertex, *Neighbors) float32
}

type Uniform float32

func (u Uniform) weigh(v Vertex, ns *Neighbors) float32 {
	return float32(u)
}

func NewWeighted(n int, p float32, w weigher) *Weighted {
	weighted := &Weighted{
		Unweighted: NewUnweighted(n, p),
		weights:    make(map[Vertex]float32, n),
	}
	for v, neighbors := range weighted.vertices {
		weighted.weights[v] = w.weigh(v, neighbors)
	}
	return weighted
}
