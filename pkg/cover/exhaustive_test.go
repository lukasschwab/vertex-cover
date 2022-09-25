package cover

import (
	"testing"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/stretchr/testify/assert"
)

func TestExhaustive(t *testing.T) {
	assert.Equal(t, float32(2), NewExhaustive(graph.GraphA).CoverWeight())
	assert.Equal(t, float32(3), NewExhaustive(graph.GraphB).CoverWeight())
	assert.Equal(t, float32(3.5), NewExhaustive(graph.GraphC).CoverWeight())
}
