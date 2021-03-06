package __1_Khapsack

import (
	"fmt"
	"testing"
)

func TestKnapsackBottomUp(t *testing.T) {
	tests := []struct {
		profits  []int
		weights  []int
		capacity int
		want     int
	}{
		{[]int{4, 5, 3, 7}, []int{2, 3, 1, 4}, 5, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v_%d", tt.weights, tt.profits, tt.capacity), func(t *testing.T) {
			if got := KnapsackBottomUpSpaceOptimized(tt.profits, tt.weights, tt.capacity); got != tt.want {
				t.Errorf("KnapsackRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
