package cover

import (
	"testing"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/stretchr/testify/assert"
)

func TestClever(t *testing.T) {
	// Always optimal for these cases.
	assert.Equal(t, float32(2), NewClever(graph.GraphA).CoverWeight())
	assert.Equal(t, float32(3), NewClever(graph.GraphB).CoverWeight())
	assert.Equal(t, float32(3.5), NewClever(graph.GraphC).CoverWeight())
}

func TestClever_Tricky(t *testing.T) {
	// Assert clever performs as Lavrov predicts.
	g := graph.NewTricky(20, 5, graph.Uniform)
	assert.Equal(t, NewClever(g).CoverWeight(), float32(25))
}
