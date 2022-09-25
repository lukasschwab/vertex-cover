package graph

import "fmt"

type Vertex int

type Neighbors struct {
	vertices map[Vertex]struct{}
}

func (ns *Neighbors) Length() int {
	return len(ns.vertices)
}

func (ns *Neighbors) Without(removed Vertex) *Neighbors {
	new := &Neighbors{
		vertices: make(map[Vertex]struct{}, len(ns.vertices)),
	}
	for vertex := range ns.vertices {
		if vertex != removed {
			new.vertices[vertex] = struct{}{}
		}
	}
	return new
}

func (ns *Neighbors) Includes(v Vertex) bool {
	_, ok := ns.vertices[v]
	return ok
}

func (ns *Neighbors) Vertices() []Vertex {
	vs := make([]Vertex, len(ns.vertices))
	i := 0
	for v := range ns.vertices {
		vs[i] = v
		i++
	}
	return vs
}

func NewNeighbors(vs []Vertex) *Neighbors {
	new := &Neighbors{
		vertices: make(map[Vertex]struct{}, len(vs)),
	}
	for _, v := range vs {
		new.vertices[v] = struct{}{}
	}
	return new
}

type Unweighted struct {
	// vertices stored as maps from the vertex to all of its neighbors; an
	// adjacency matrix.
	vertices map[Vertex]*Neighbors
}

func (u *Unweighted) Vertices() []Vertex {
	vertices := make([]Vertex, len(u.vertices))
	i := 0
	for k := range u.vertices {
		vertices[i] = k
		i++
	}
	return vertices
}

func (u *Unweighted) Degree(v Vertex) int {
	return u.vertices[v].Length()
}

func (u *Unweighted) Neighbors(v Vertex) *Neighbors {
	return u.vertices[v]
}

func (u *Unweighted) Print() {
	fmt.Printf("Graph with size %v:\n", len(u.vertices))
	for v, neighbors := range u.vertices {
		fmt.Printf("\t%v: %v (%v)\n", v, neighbors.vertices, u.Degree(v))
	}
}

// NOTE: this would be faster if it were destructive.
func (u *Unweighted) Without(removed Vertex) *Unweighted {
	new := &Unweighted{
		vertices: make(map[Vertex]*Neighbors, len(u.vertices)-1),
	}
	for vertex, neighbors := range u.vertices {
		if vertex != removed {
			new.vertices[vertex] = neighbors.Without(removed)
		}
	}
	return new
}

type Weighted struct {
	*Unweighted
	weights map[Vertex]float32
}

func (w *Weighted) Weight(v Vertex) float32 {
	return w.weights[v]
}

func (w *Weighted) Without(removed Vertex) *Weighted {
	new := &Weighted{
		Unweighted: w.Unweighted.Without(removed),
		weights:    make(map[Vertex]float32, len(w.weights)-1),
	}
	for vertex, weight := range w.weights {
		if vertex != removed {
			new.weights[vertex] = weight
		}
	}
	return new
}

func (w *Weighted) Print() {
	w.Unweighted.Print()
	fmt.Printf("Weights: %v\n", w.weights)
}
