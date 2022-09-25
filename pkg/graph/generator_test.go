package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnweighted(t *testing.T) {
	n := 10
	g := NewUnweighted(n, 0.5)
	assert.Len(t, g.Vertices(), n)

	g.Print()
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

	g.Print()
}
