package _322_coin_change

import "math"

// todo 1 read and understand bottom to up (approach 3)

func coinChangeBruteForce(coins []int, amount int) int {
	return bruteForceHelper(0, coins, amount)
}

func bruteForceHelper(idxCoin int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if idxCoin < len(coins) {
		maxVal := amount / coins[idxCoin]
		minCost := math.MaxInt64
		for x := 0; x <= maxVal; x++ {
			if amount >= x*coins[idxCoin] {
				res := bruteForceHelper(idxCoin+1, coins, amount-x*coins[idxCoin])
				if res != -1 {
					minCost = min(minCost, res+x)
				}
			}
		}
		if minCost == math.MaxInt64 {
			return -1
		}
		return minCost
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
