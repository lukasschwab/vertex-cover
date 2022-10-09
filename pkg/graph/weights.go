package graph

import (
	"math/rand"
)

// Named weighers.
var (
	Uniform        Weigher = uniform{}
	Random                 = random{}
	DegreeNegative         = degreeNegative{}
	DegreePositive         = degreePositive{}
)

// Weigher for vertices in a [graph.Weighted].
type Weigher interface {
	weigh(Vertex, *Neighbors) float32
}

// Uniform Weigher. Every vertex gets weight 1.
type uniform struct{}

func (u uniform) weigh(v Vertex, ns *Neighbors) float32 {
	return float32(1)
}

// Random Weigher. Every vertex gets a pseudorandom weight on [0, 1).
type random struct{}

func (r random) weigh(v Vertex, ns *Neighbors) float32 {
	return rand.Float32()
}

// DegreeNegative Weigher gives each vertex a weight inversely proportional to
// its degree.
type degreeNegative struct{}

func (d degreeNegative) weigh(v Vertex, ns *Neighbors) float32 {
	return 1.0 / (float32(ns.Length()) - 0.0001)
}

// DegreePositive gives each vertex a weight equal to its degree.
type degreePositive struct{}

func (d degreePositive) weigh(v Vertex, ns *Neighbors) float32 {
	return float32(ns.Length())
}
