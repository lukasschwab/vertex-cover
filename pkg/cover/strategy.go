package cover

type Strategy interface {
	CoverWeight() float32
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
