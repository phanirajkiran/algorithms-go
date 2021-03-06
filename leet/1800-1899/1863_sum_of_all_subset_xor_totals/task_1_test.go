package _863_sum_of_all_subset_xor_totals

import (
	"fmt"
	"testing"
)

func Test_subsetXORSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3}, 6},
		{[]int{5, 1, 6}, 28},
		{[]int{3, 4, 5, 6, 7, 8}, 480},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := subsetXORSum(tt.nums); got != tt.want {
				t.Errorf("subsetXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
