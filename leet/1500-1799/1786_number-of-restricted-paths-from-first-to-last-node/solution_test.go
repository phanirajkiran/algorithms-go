package _1786_number_of_restricted_paths_from_first_to_last_node

import (
	"testing"
)

func Test_countRestrictedPaths(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"1", 5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}, 3},
		{"2", 7, [][]int{{1, 3, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}, {2, 6, 4}}, 1},
		{"312 size input", 312, bigArrOfEdges, 926335851},
		{"422 size input", 422, bigArrOfEdges_422, 13944179},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRestrictedPaths3(tt.n, tt.edges); got != tt.want {
				t.Errorf("countRestrictedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
