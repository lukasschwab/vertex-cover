package graph

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Vertex in a graph. The int representation is an ID rather than a weight.
type Vertex int

// String name for v.
func (v Vertex) String() string {
	return fmt.Sprintf("%d", v)
}

// Neighbors of some vertex in a graph: a set of other vertices in the graph.
//
// TODO: this currently implements a directed graph, but it's really only meant
// for use with unidrected graphs (enforced by generator_test.go). Could
// refactor this neighbors model to enforce undirectedness.
type Neighbors struct {
	vertices map[Vertex]struct{}
}

// Length of the list of neighbor vertices ns.
func (ns *Neighbors) Length() int {
	return len(ns.vertices)
}

// Without returns ns minus a removed vertex.
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

// Includes returns whether v is among the vertices in ns.
func (ns *Neighbors) Includes(v Vertex) bool {
	_, ok := ns.vertices[v]
	return ok
}

// Vertices (neighbors) in ns.
func (ns *Neighbors) Vertices() []Vertex {
	vs := make([]Vertex, len(ns.vertices))
	i := 0
	for v := range ns.vertices {
		vs[i] = v
		i++
	}
	return vs
}

// NewNeighbors returns an empty [Neighbors] set.
func NewNeighbors(vs []Vertex) *Neighbors {
	new := &Neighbors{
		vertices: make(map[Vertex]struct{}, len(vs)),
	}
	for _, v := range vs {
		new.vertices[v] = struct{}{}
	}
	return new
}

// Unweighted graph.
type Unweighted struct {
	// vertices stored as maps from the vertex to all of its neighbors; an
	// adjacency matrix.
	vertices map[Vertex]*Neighbors
}

// Edges in u.
func (u *Unweighted) Edges() [][2]Vertex {
	edges := [][2]Vertex{}
	for _, v := range u.Vertices() {
		for _, neighbor := range u.Neighbors(v).Vertices() {
			if neighbor >= v {
				edges = append(edges, [2]Vertex{v, neighbor})
			}
		}
	}
	return edges
}

// Vertices of u.
func (u *Unweighted) Vertices() []Vertex {
	vertices := make([]Vertex, len(u.vertices))
	i := 0
	for k := range u.vertices {
		vertices[i] = k
		i++
	}
	return vertices
}

// Add a vertex to u.
func (u *Unweighted) Add(a Vertex) {
	u.vertices[a] = NewNeighbors([]Vertex{})
}

// Connect vertices a, b in u.
func (u *Unweighted) Connect(a, b Vertex) {
	u.vertices[a].vertices[b] = struct{}{}
	u.vertices[b].vertices[a] = struct{}{}
}

// Degree of v in u.
func (u *Unweighted) Degree(v Vertex) int {
	return u.vertices[v].Length()
}

// Neighbors of v in u.
func (u *Unweighted) Neighbors(v Vertex) *Neighbors {
	return u.vertices[v]
}

// Print a summary of u.
func (u *Unweighted) Print() {
	fmt.Printf("Graph with size %v:\n", len(u.vertices))
	for v, neighbors := range u.vertices {
		fmt.Printf("\t%v: %v (%v)\n", v, neighbors.vertices, u.Degree(v))
	}
}

// Display based on https://github.com/go-echarts/examples/blob/master/examples/graph.go
//
// TODO: move this onto Weighted and display the vertex weights.
func (u *Unweighted) Display(title string) *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
	)

	nodes := make([]opts.GraphNode, len(u.Vertices()))
	links := []opts.GraphLink{}
	for i, v := range u.Vertices() {
		nodes[i] = opts.GraphNode{Name: fmt.Sprintf("%d", v)}
		for _, neighbor := range u.Neighbors(v).Vertices() {
			// TODO: include edge weights for weighted graphs.
			links = append(links, opts.GraphLink{Source: v.String(), Target: neighbor.String()})
		}
	}

	graph.AddSeries("graph", nodes, links).SetSeriesOptions(
		charts.WithGraphChartOpts(opts.GraphChart{
			Force:  &opts.GraphForce{Repulsion: 8000},
			Layout: "circular",
		}),
		charts.WithLabelOpts(opts.Label{Show: true, Position: "right"}),
	)

	return graph
}

// Without returns u minus a removed vertex, its weights, and its incident
// edges.
//
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

// Weighted graph, with weighted vertices.
type Weighted struct {
	*Unweighted
	weights map[Vertex]float32
}

// Weight of a vertex in w.
func (w *Weighted) Weight(v Vertex) float32 {
	return w.weights[v]
}

// Without returns w minus a removed vertex, its weights, and its incident
// edges.
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

// Print a summary of w.
func (w *Weighted) Print() {
	w.Unweighted.Print()
	fmt.Printf("Weights: %v\n", w.weights)
}
