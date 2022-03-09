package bfs

import (
	"algorithm-recipe-golang/pkg/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFSUndirected_CountEdgeAndVertices(t *testing.T) {
	g := utils.Graph1(false)

	var discoveredVertexCount, discoveredEdgeCount int
	processVertexEarly := func(v int) {
		fmt.Printf("v: %d\n", v)
		discoveredVertexCount++
	}
	processVertexLate := func(v int) {}
	processEdge := func(u, v int) {
		fmt.Printf("u -> v: %d -> %d\n", u, v)
		discoveredEdgeCount++
	}

	err := BFS(g, 0, processVertexEarly, processVertexLate, processEdge)

	assert.NoError(t, err)
	assert.Equal(t, 5, discoveredVertexCount)
	assert.Equal(t, 5, discoveredEdgeCount)
}
