package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnweighted(t *testing.T) {
	n := 10
	g := NewUnweighted(n, 0.5)
	assert.Len(t, g.Vertices(), n)

	// Undirected graph.
	for _, v := range g.Vertices() {
		for other := range g.Neighbors(v).vertices {
			assert.Contains(t, g.Neighbors(other).vertices, v)
		}
	}
}

func TestNewWeighted(t *testing.T) {
	n := 10
	w := Uniform(0.333)
	g := NewWeighted(n, 0.5, w)
	assert.Len(t, g.Vertices(), n)
	assert.Len(t, g.weights, n)
	for _, weight := range g.weights {
		assert.Equal(t, weight, float32(w))
	}
	// Undirected graph.
	for _, v := range g.Vertices() {
		for _, other := range g.Neighbors(v).Vertices() {
			assert.Contains(t, g.Neighbors(other).Vertices(), v)
		}
	}
}
