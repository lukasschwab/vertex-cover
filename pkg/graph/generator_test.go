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
	g := NewWeighted(n, 0.5, Uniform{})
	assert.Len(t, g.Vertices(), n)
	assert.Len(t, g.weights, n)
	for _, weight := range g.weights {
		assert.Equal(t, weight, 1)
	}
	// Undirected graph.
	for _, v := range g.Vertices() {
		for _, other := range g.Neighbors(v).Vertices() {
			assert.Contains(t, g.Neighbors(other).Vertices(), v)
		}
	}
}

func TestNewTricky(t *testing.T) {
	// Test the example graph: https://faculty.math.illinois.edu/~mlavrov/docs/482-spring-2020/lecture36.pdf
	n := 20
	k := 5
	g := NewTricky(n, k, Uniform{})
	assert.Len(t, g.Vertices(), 45)
	// A vertices.
	for v := Vertex(0); v < Vertex(20); v++ {
		degree := g.Degree(v)
		assert.True(t, degree == 3 || degree == 4)
	}
	// B2.
	for v := Vertex(20); v < Vertex(30); v++ {
		degree := g.Degree(v)
		assert.True(t, degree == 2 || degree == 3)
	}
	// B3.
	for v := Vertex(30); v < Vertex(36); v++ {
		degree := g.Degree(v)
		assert.True(t, degree == 3 || degree == 4)
	}
	// B4.
	for v := Vertex(36); v < Vertex(41); v++ {
		degree := g.Degree(v)
		assert.True(t, degree == 4 || degree == 5)
	}
	// B5.
	for v := Vertex(41); v < Vertex(45); v++ {
		degree := g.Degree(v)
		assert.True(t, degree == 5 || degree == 6)
	}
	assert.Len(t, g.Edges(), n*(k-1))
}
