package cover

// Strategy (algorithm) for finding a small vertex cover.
type Strategy interface {
	CoverWeight() float32
}

// Weight of a candidate vertex cover. noCover indicates the candidate isn't a
// vertex cover on its graph.
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
