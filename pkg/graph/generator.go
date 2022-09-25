package graph

import "math/rand"

func NewUnweighted(n int, p float32) *Unweighted {
	u := &Unweighted{
		vertices: make(map[Vertex]*Neighbors, n),
	}

	for i := Vertex(0); int(i) < n; i++ {
		neighbors := &Neighbors{
			vertices: make(map[Vertex]struct{}),
		}
		for candidate := Vertex(0); int(candidate) < n; candidate++ {
			if rand.Float32() < p {
				neighbors.vertices[candidate] = struct{}{}
			}
		}
		u.vertices[i] = neighbors
	}

	return u
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
