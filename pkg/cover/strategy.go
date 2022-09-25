package cover

import "github.com/lukasschwab/vertex-cover/pkg/graph"

type Strategy interface {
	CoverWeight(w *graph.Weighted) float32
}

type Weight struct {
	float32
	noCover bool
}

// OrLesser returns the lesser of w and other.
func (w Weight) OrLesser(other Weight) Weight {
	if w.noCover {
		return other
	}
	if other.noCover {
		return w
	}

	if w.float32 <= other.float32 {
		return w
	}
	return other
}
