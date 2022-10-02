package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var f struct{}

func TestNeighbors_Length(t *testing.T) {
	n := &Neighbors{
		vertices: map[Vertex]struct{}{1: f, 2: f, 3: f},
	}
	assert.Equal(t, 3, n.Length())
	n.vertices[4] = f
	assert.Equal(t, 4, n.Length())
}

func TestNeighbors_Without(t *testing.T) {
	n := &Neighbors{
		vertices: map[Vertex]struct{}{1: f, 2: f, 3: f},
	}
	var removed Vertex = 3
	nprime := n.Without(removed)
	assert.Len(t, nprime.vertices, 2)
	assert.NotContains(t, nprime.vertices, removed)
}

func TestVertex_String(t *testing.T) {
	assert.Equal(t, "0", Vertex(0).String())
	assert.Equal(t, "18", Vertex(18).String())
}
