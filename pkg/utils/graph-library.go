package utils

import "algorithm-recipe-golang/pkg"

// Graph1 sample graph for testing
// (0) <--- (1) (4)
//  |      / |   ^
//  |     /  |  /
//  v    /   v /
// (2)<-.   (3)
func Graph1(directed bool) *pkg.AdjList {
	g := pkg.NewAdjList(5, directed)

	g.AddEdge(0, 2)
	g.AddEdge(1, 0)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)

	return g
}
