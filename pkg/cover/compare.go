package cover

import "github.com/lukasschwab/vertex-cover/pkg/graph"

// StrategyTemplate is a helper type, since the strategy constructors take
// graphs. It might make more sense to have the strategy interface itself take
// a graph.
type StrategyTemplate func(*graph.Weighted) Strategy

// Comparison between baseline and test strategies.
type Comparison struct {
	Baseline, Test StrategyTemplate
}

// Result of a comparison between a baseline strategy and a test strategy on a
// graph.
type Result struct {
	Baseline, Test float32
}

// Delta in weight between the comparison's baseline result and the test result.
func (r Result) Delta() float32 {
	return r.Test - r.Baseline
}

// Run a comparison on g.
func (c Comparison) Run(g *graph.Weighted) Result {
	return Result{
		Baseline: c.Baseline(g).CoverWeight(),
		Test:     c.Test(g).CoverWeight(),
	}
}
