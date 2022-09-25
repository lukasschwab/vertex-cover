package cover

import (
	"testing"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/stretchr/testify/assert"
)

func TestVazirani(t *testing.T) {
	assert.Equal(
		t,
		float32(6),
		NewVazirani(graph.GraphVaziraniFail).CoverWeight(),
		"Known-bad Vazirani graph: http://lukasschwab.me/blog/gen/graphs-at-work.html#fn7",
	)
}
