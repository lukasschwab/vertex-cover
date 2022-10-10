package graph

import (
	"math/rand"
)

// Named weighers.
var (
	// Uniform Weigher. Every vertex gets weight 1.
	Uniform Weigher = uniform{}
	// Random Weigher. Every vertex gets a pseudorandom weight on [0, 1).
	Random = random{}
	// DegreeNegative gives each vertex a weight inversely proportional to its
	// degree.
	DegreeNegative = degreeNegative{}
	// DegreePositive gives each vertex a weight equal to its degree.
	DegreePositive = degreePositive{}
	// DegreePositiveSuperlinear gives each vertex a weight equal to the square
	// of its degree.
	DegreePositiveSuperlinear = degreePositiveSuperlinear{}
)

// Weigher for vertices in a [graph.Weighted].
type Weigher interface {
	weigh(Vertex, *Neighbors) float32
}

type uniform struct{}

func (u uniform) weigh(v Vertex, ns *Neighbors) float32 {
	return float32(1)
}

type random struct{}

func (r random) weigh(v Vertex, ns *Neighbors) float32 {
	return rand.Float32()
}

type degreeNegative struct{}

func (d degreeNegative) weigh(v Vertex, ns *Neighbors) float32 {
	return 1.0 / (float32(ns.Length()) - 0.0001)
}

type degreePositive struct{}

func (d degreePositive) weigh(v Vertex, ns *Neighbors) float32 {
	return float32(ns.Length()) + 0.1
}

type degreePositiveSuperlinear struct{}

func (d degreePositiveSuperlinear) weigh(v Vertex, ns *Neighbors) float32 {
	return float32(ns.Length()*ns.Length()) + 0.1
}
