package graph

var (
	a = Vertex(1)
	b = Vertex(2)
	c = Vertex(3)
)

// GraphA is a test graph described in "Graphs at Work."
// Minimal vertex cover: b.
var GraphA = &Weighted{
	Unweighted: &Unweighted{
		vertices: map[Vertex]*Neighbors{
			a: NewNeighbors([]Vertex{b}),
			b: NewNeighbors([]Vertex{a, c}),
			c: NewNeighbors([]Vertex{b}),
		},
	},
	weights: map[Vertex]float32{a: 1.5, b: 2, c: 1.5},
}

// GraphB is a test graph described in "Graphs at Work."
// Minimal vertex cover: a and c.
var GraphB = &Weighted{
	Unweighted: &Unweighted{
		vertices: map[Vertex]*Neighbors{
			a: NewNeighbors([]Vertex{b, a}),
			b: NewNeighbors([]Vertex{a, c}),
			c: NewNeighbors([]Vertex{b}),
		},
	},
	weights: map[Vertex]float32{a: 1.5, b: 2, c: 1.5},
}

// GraphC is a test graph described in "Graphs at Work."
// Optimal: a and b.
var GraphC = &Weighted{
	Unweighted: &Unweighted{
		vertices: map[Vertex]*Neighbors{
			a: NewNeighbors([]Vertex{b, a}),
			b: NewNeighbors([]Vertex{a, c}),
			c: NewNeighbors([]Vertex{b}),
		},
	},
	weights: map[Vertex]float32{a: 1.5, b: 2, c: 3},
}

// GraphVaziraniFail is a test graph described in "Graphs at Work," for which
// the Vazirani algorithm performs poorly.
var GraphVaziraniFail = &Weighted{
	Unweighted: &Unweighted{
		vertices: map[Vertex]*Neighbors{
			a: NewNeighbors([]Vertex{b}),
			b: NewNeighbors([]Vertex{a, c}),
			c: NewNeighbors([]Vertex{b}),
		},
	},
	weights: map[Vertex]float32{a: 1.5, b: 3, c: 1.5},
}
