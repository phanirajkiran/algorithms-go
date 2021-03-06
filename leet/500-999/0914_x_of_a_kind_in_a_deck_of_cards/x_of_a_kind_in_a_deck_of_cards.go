package _914_x_of_a_kind_in_a_deck_of_cards

import (
	"sort"
)

// passed. best solution. tptl
func hasGroupsSizeGCD(deck []int) bool {
	m := make(map[int]int)
	for _, v := range deck {
		m[v]++
	}

	var g int
	for _, v := range m {
		if g = gcd(g, v); g < 2 {
			return false
		}
	}

	return true
}

// passed
func hasGroupsSizeXBruteForce(deck []int) bool {
	n := len(deck)
	count := make([]int, 10000)
	for _, num := range deck {
		count[num]++
	}

	var values []int
	for i := 0; i < 10000; i++ {
		if count[i] > 0 {
			values = append(values, count[i])
		}
	}

Loop:
	for X := 2; X <= n; X++ {
		if n%X == 0 {
			for _, v := range values {
				if v%X != 0 {
					continue Loop
				}
			}

			return true
		}
	}

	return false
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}

	return gcd(y%x, x)
}

// passed
func hasGroupsSizeX(deck []int) bool {
	n := len(deck)
	sort.Ints(deck)
	currLen := 1

	var divisor int
	for i := 1; i <= n; i++ {
		if i < n && deck[i] == deck[i-1] {
			currLen++
			continue
		}

		if divisor = gcd(divisor, currLen); divisor < 2 {
			return false
		}
		currLen = 1
	}

	return true
}
