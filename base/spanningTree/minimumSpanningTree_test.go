package spanTree

import (
	"github.com/babadro/algorithms-go/base/graph"
	"reflect"
	"testing"
)

func TestKruskalMST(t *testing.T) {
	tests := []struct {
		name          string
		vertexCount   int
		edges         [][3]int
		wantWeightSum int
	}{
		{
			name:        "1",
			vertexCount: 9,
			edges: [][3]int{
				{0, 1, 4},
				{0, 7, 8},
				{1, 7, 11},
				{1, 2, 8},
				{7, 8, 7},
				{7, 6, 1},
				{2, 8, 2},
				{8, 6, 6},
				{2, 3, 7},
				{2, 5, 4},
				{6, 5, 2},
				{3, 5, 14},
				{3, 4, 9},
				{5, 4, 10},
			},
			wantWeightSum: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input []Edge
			for _, e := range tt.edges {
				input = append(input, arrayToEdge(e))
			}

			gotMst := KruskalMST(tt.vertexCount, input)

			sum := 0
			for _, e := range gotMst {
				sum += e.weight
			}

			g := graph.New(tt.vertexCount)
			for _, edge := range gotMst {
				g.AddEdge(edge.src, edge.dst)
				g.AddEdge(edge.dst, edge.src)
			}

			visited := make([]bool, tt.vertexCount)
			f := func(v int) {
				visited[v] = true
			}
			g.BFS(0, f)
			for vertex, ok := range visited {
				if !ok {
					t.Errorf("there is unconnected vertex=%d in result spanning tree", vertex)
				}
			}

			if !reflect.DeepEqual(sum, tt.wantWeightSum) {
				t.Errorf("KruskalMST() = %v, want %v", sum, tt.wantWeightSum)
			}
		})
	}
}

func arrayToEdge(arr [3]int) Edge {
	return Edge{
		src:    arr[0],
		dst:    arr[1],
		weight: arr[2],
	}
}

func TestPrimMST(t *testing.T) {
	tests := []struct {
		name          string
		vertexCount   int
		edges         [][3]int
		wantWeightSum int
	}{
		{
			name:        "1",
			vertexCount: 5,
			edges: [][3]int{
				{0, 1, 2},
				{0, 3, 6},
				{1, 4, 5},
				{1, 3, 8},
				{1, 2, 3},
				{2, 4, 7},
				{3, 4, 9},
			},
			wantWeightSum: 16,
		},
		{
			name:        "2",
			vertexCount: 9,
			edges: [][3]int{
				{0, 1, 4},
				{0, 7, 8},
				{1, 7, 11},
				{1, 2, 8},
				{7, 8, 7},
				{7, 6, 1},
				{2, 8, 2},
				{8, 6, 6},
				{2, 3, 7},
				{2, 5, 4},
				{6, 5, 2},
				{3, 5, 14},
				{3, 4, 9},
				{5, 4, 10},
			},
			wantWeightSum: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adjMatrix := graph.EdgesToAdjMatrix(tt.vertexCount, tt.edges)
			gotMst := PrimMST(tt.vertexCount, adjMatrix)

			sum := 0
			for _, e := range gotMst {
				sum += e.weight
			}

			g := graph.New(tt.vertexCount)
			for _, edge := range gotMst {
				g.AddEdge(edge.src, edge.dst)
				g.AddEdge(edge.dst, edge.src)
			}

			visited := make([]bool, tt.vertexCount)
			f := func(v int) {
				visited[v] = true
			}
			g.BFS(0, f)
			for vertex, ok := range visited {
				if !ok {
					t.Errorf("there is unconnected vertex=%d in result spanning tree", vertex)
				}
			}

			if !reflect.DeepEqual(sum, tt.wantWeightSum) {
				t.Errorf("PrimMST() = %v, want %v", sum, tt.wantWeightSum)
			}
		})
	}
}
