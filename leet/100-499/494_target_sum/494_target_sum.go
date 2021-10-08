package _494_target_sum

func findTargetSumWaysBruteForce(nums []int, target int) int {
	return recBruteForce(nums, 0, 0, target)
}

func recBruteForce(nums []int, idx, cur, target int) int {
	if idx == len(nums) {
		if cur == target {
			return 1
		}

		return 0
	}

	res1 := recBruteForce(nums, idx+1, cur+nums[idx], target)
	res2 := recBruteForce(nums, idx+1, cur-nums[idx], target)

	return res1 + res2
}

// tptl. passed. #medium (hard for me). #dp
func findTargetSumTopDown(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total < target || target < -total || (total+target)%2 == 1 {
		return 0
	}

	sum := (total + target) / 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, nums, 0, sum)
}

func recTopDown(dp [][]int, nums []int, idx, sum int) int {
	if idx == len(nums) {
		if sum == 0 {
			return 1
		}

		return 0
	}

	if sum < 0 {
		return 0
	}

	if dp[idx][sum] == -1 {
		res1 := recTopDown(dp, nums, idx+1, sum-nums[idx])
		res2 := recTopDown(dp, nums, idx+1, sum)

		dp[idx][sum] = res1 + res2
	}

	return dp[idx][sum]
}