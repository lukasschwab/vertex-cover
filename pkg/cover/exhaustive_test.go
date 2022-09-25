package cover

import (
	"testing"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/stretchr/testify/assert"
)

func TestExhaustive(t *testing.T) {
	graph.GraphA.Print()
	eA := Exhaustive{graph.GraphA}
	assert.Equal(t, float32(2), eA.CoverWeight())

	graph.GraphB.Print()
	eB := Exhaustive{graph.GraphB}
	assert.Equal(t, float32(3), eB.CoverWeight())

	graph.GraphC.Print()
	eC := Exhaustive{graph.GraphC}
	assert.Equal(t, float32(3.5), eC.CoverWeight())
}
