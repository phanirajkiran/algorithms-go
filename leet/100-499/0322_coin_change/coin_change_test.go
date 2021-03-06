package _322_coin_change

import "testing"

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{"1", []int{1, 2, 5}, 11, 3},
		{"2", []int{2}, 3, -1},
		{"3", []int{1}, 0, 0},
		{"4", []int{1}, 1, 1},
		{"5", []int{1}, 2, 2},
		{"6", []int{186, 419, 83, 408}, 6249, 20},
		{"7", []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangeTopDown2(tt.coins, tt.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
