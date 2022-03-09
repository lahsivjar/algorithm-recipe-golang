package pkg

import (
	"errors"
	"fmt"
)

// AdjList models adjacency list representation of a graph
type AdjList struct {
	graph    []*Node
	directed bool
}

// NewAdjList creates a new adjacency list for a given number of vertices
func NewAdjList(vertices int, directed bool) *AdjList {
	return &AdjList{
		graph:    make([]*Node, vertices),
		directed: directed,
	}
}

// Print prints the graph for visualization
func (a *AdjList) Print() {
	for i := 0; i < len(a.graph); i++ {
		fmt.Printf("%d", i)
		head := a.graph[i]

		for head != nil {
			fmt.Printf(" --> %d", head.Value)
			head = head.Next
		}
		fmt.Printf(" --> nil\n")
	}
}

// Size returns number of vertices in the graph
func (a *AdjList) Size() int {
	return len(a.graph)
}

// IsDirected returns true for directed graphs, false otherwise
func (a *AdjList) IsDirected() bool {
	return a.directed
}

// GetAdjVertices returns the head of the linked list for adjacency vertices of the given vertex
func (a *AdjList) GetAdjVertices(v int) (*Node, error) {
	if v >= len(a.graph) {
		return nil, errors.New("vertex out of range")
	}
	return a.graph[v], nil
}

// GetAdjVerticesList returns a list of adjacency vertices for the given vertex
func (a *AdjList) GetAdjVerticesList(v int) ([]int, error) {
	head, err := a.GetAdjVertices(v)
	if err != nil {
		return nil, err
	}

	var result []int
	for head != nil {
		result = append(result, head.Value)
		head = head.Next
	}
	return result, nil
}

// AddEdge adds a new edge in the graph
func (a *AdjList) AddEdge(from, to int) error {
	if from >= len(a.graph) || to >= len(a.graph) {
		return errors.New("one of the provided vertex is out of range")
	}

	headFrom := a.graph[from]
	a.graph[from] = &Node{
		Value: to,
		Next:  headFrom,
	}

	if !a.directed {
		headTo := a.graph[to]
		a.graph[to] = &Node{
			Value: from,
			Next:  headTo,
		}
	}

	return nil
}
