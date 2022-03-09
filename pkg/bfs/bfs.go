package bfs

import (
	"algorithm-recipe-golang/pkg"
	"container/list"
)

type graphRep interface {
	IsDirected() bool
	Size() int
	GetAdjVertices(int) (*pkg.Node, error)
}

type nodeStatus int

const (
	undiscovered nodeStatus = iota
	discovered
	processed
)

// BFS performs a breadth first search over a graph
func BFS(
	g graphRep,
	root int,
	processVertexEarly, processVertexLate func(v int), // uniquely process vertex at different states
	processEdge func(from, to int), // uniquely process edge
) error {
	n := g.Size()
	// Initalize them with undiscovered
	status := make([]nodeStatus, n)

	// Create a queue to track what to process next
	q := list.New()
	q.PushBack(root)
	status[root] = discovered

	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		// We can process the vertex here; before discovering all it's edges
		processVertexEarly(u)

		head, err := g.GetAdjVertices(u)
		if err != nil {
			return err
		}

		for head != nil {
			v := head.Value
			// Here we will discover all edges u->v
			// If g is un-directed and v is procssed then u--v has already been procssed as v--u
			// If g is directed then all edges are new edges since even if v is processed it will
			// only process for v->u, not u->v and both of these are seperate edges for directed
			if status[v] != processed || g.IsDirected() {
				// We can process the new edge u->v
				processEdge(u, v)
			}

			// We will only push the undiscovered node to v since discovered nodes are already
			// scheduled to be processed and processed nodes are already processed
			if status[v] == undiscovered {
				status[v] = discovered
				q.PushBack(v)
			}

			head = head.Next
		}

		// We can process the node here; after discovering all it's neighbours
		processVertexLate(u)
		status[u] = processed
	}

	return nil
}
