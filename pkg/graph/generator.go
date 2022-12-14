package graph

import (
	"math"
	"math/rand"
)

// NewUnweighted graph with n vertices, where any two vertices are adjacent with
// probability p in [0, 1).
func NewUnweighted(n int, p float32) *Unweighted {
	g := &Unweighted{
		vertices: make(map[Vertex]*Neighbors, n),
	}
	// Initiate vertices, empty neighbors.
	for i := Vertex(0); i < Vertex(n); i++ {
		g.vertices[i] = NewNeighbors([]Vertex{})
	}

	vertices := g.Vertices()
	for i, u := range vertices {
		// Disallow u<>u edges: they force u's inclusion, whcih forces
		// weight to n as p approaches 1.
		for j := i + 1; j < len(vertices); j++ {
			v := vertices[j]
			if rand.Float32() < p {
				// TODO: refactor. Blech.
				g.vertices[u].vertices[v] = struct{}{}
				g.vertices[v].vertices[u] = struct{}{}
			}
		}
	}

	return g
}

// NewWeighted graph with w-determined weights bolted onto [NewUnweighted].
func NewWeighted(n int, p float32, w Weigher) *Weighted {
	weighted := &Weighted{
		Unweighted: NewUnweighted(n, p),
		weights:    make(map[Vertex]float32, n),
	}
	for v, neighbors := range weighted.vertices {
		weighted.weights[v] = w.weigh(v, neighbors)
	}
	return weighted
}

type vertexGenerator struct {
	last Vertex
}

func (g *vertexGenerator) next() Vertex {
	g.last++
	return g.last
}

// NewTricky bimodal graph designed to produce very suboptimal outcomes for the
// clever solution.
//
// - a: the number of vertices in the true minimal set a.
// - k: the number of b-groups. k ≤ a.
//
// The resulting graph has n vertices where n = a * Hk, for the kth harmonic Hk.
// https://en.wikipedia.org/wiki/Harmonic_number
//
// Lavrov describes this construction on page 2:
// https://faculty.math.illinois.edu/~mlavrov/docs/482-spring-2020/lecture36.pdf
func NewTricky(a, k int, w Weigher) *Weighted {
	graph := &Unweighted{vertices: make(map[Vertex]*Neighbors)}

	// https://faculty.math.illinois.edu/~mlavrov/docs/482-spring-2020/lecture36.pdf

	g := &vertexGenerator{-1}

	A := make([]Vertex, a)
	for i := range A {
		A[i] = g.next()
		graph.Add(A[i])
	}

	// Bs := make([][]Vertex, k-1)
	for i := 2; i <= k; i++ {
		Ai := 0
		// Construct the ith B-set.
		Bi := make([]Vertex, int(math.Floor(float64(a)/float64(i))))
		for j := range Bi {
			Bi[j] = g.next()
			graph.Add(Bi[j])
			for conns := 0; conns < i; conns++ {
				graph.Connect(Bi[j], A[Ai])
				Ai++
			}
		}
		// Divvy up remaining connections to A.
		bi := 0
		for Ai < len(A) && bi < len(Bi) {
			// TODO: Connect Bi[bi] to A[Ai].
			graph.Connect(Bi[bi], A[Ai])
			Ai++
			bi++
		}
	}

	weighted := &Weighted{
		Unweighted: graph,
		weights:    make(map[Vertex]float32, len(graph.vertices)),
	}
	for v, neighbors := range graph.vertices {
		weighted.weights[v] = w.weigh(v, neighbors)
	}
	return weighted
}
