package _38

import (
	"github.com/babadro/algorithms-go/utils"
	"sort"
)

// todo 1
func maxBuilding(n int, restrictions [][]int) int {
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	rLen := len(restrictions)
	maxH := 0
	lastX, lastY := 1, 0
	for i := 0; i <= rLen; i++ {
		var r []int
		if i == rLen {
			if (rLen > 0 && restrictions[rLen-1][0] < n) || rLen == 0 {
				r = []int{n, 1_000_000_000}
			} else {
				break
			}
		} else {
			r = restrictions[i]
		}

		x, y := r[0], r[1]
		diffX := x - lastX
		diffY := utils.Abs(lastY - y)
		minY := utils.Min(lastY, y)

		currHeight := height(diffX, diffY, minY)
		maxH = utils.Max(maxH, currHeight)

		lastX, lastY = x, y
	}

	return maxH
}

func height(diffX, diffY, minY int) int {
	diffY = utils.Min(diffY, diffX)
	idealDiff := diffX

	d1 := idealDiff - diffY

	return minY + idealDiff - (d1/2 + d1%2)
}
