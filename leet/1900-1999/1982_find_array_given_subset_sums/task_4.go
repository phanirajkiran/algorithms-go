package _982_find_array_given_subset_sums

import (
	"math"
	"sort"
)

// tle
func recoverArrayTle(n int, sums []int) []int {
	sort.Ints(sums)

	var res []int
	req(sums, make([]int, 0, n), &res, 0, n)

	return res
}

func req(sums, curr []int, res *[]int, idx, n int) {
	if n == 0 {
		candidate := getSums(curr)
		sort.Ints(candidate)
		if intSlicesAreEqual(sums, candidate) {
			ans := make([]int, len(curr))
			copy(ans, curr)
			*res = ans
		}

		return
	}

	if idx == len(sums) || n < 0 {
		return
	}

	req(sums, curr, res, idx+1, n)
	req(sums, append(curr, sums[idx]), res, idx+1, n-1)
}

func getSums(nums []int) []int {
	n := len(nums)
	lenSums := int(math.Pow(float64(n), 2))
	res, subArr := make([]int, 0, lenSums), make([]int, 0, n)
	reqGetSums(nums, subArr, &res, 0)

	return res
}

func reqGetSums(nums, curr []int, res *[]int, idx int) {
	if idx == len(nums) {
		sum := 0
		for _, num := range curr {
			sum += num
		}

		*res = append(*res, sum)

		return
	}

	reqGetSums(nums, curr, res, idx+1)
	reqGetSums(nums, append(curr, nums[idx]), res, idx+1)
}

func intSlicesAreEqual(slice1, slice2 []int) bool {
	length := len(slice1)
	if length != len(slice2) {
		return false
	}
	for i := 0; i < length; i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
