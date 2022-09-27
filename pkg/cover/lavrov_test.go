package cover

import (
	"testing"

	"github.com/lukasschwab/vertex-cover/pkg/graph"
	"github.com/stretchr/testify/assert"
)

func TestLavrov(t *testing.T) {
	for a := 20; a <= 100; a++ {
		g := graph.NewTricky(a, 17, graph.Uniform(1))
		assert.LessOrEqual(
			t, NewLavrov(g).CoverWeight(), float32(2*a),
			"Lavrov's greedy algorithm is a 2-approximation",
		)
	}
}
