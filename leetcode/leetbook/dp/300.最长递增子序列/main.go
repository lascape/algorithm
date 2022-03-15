package main

import "math"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp = make([]int, len(nums))
	dp[0] = 1
	maxans := 1
	for i, a := range nums {
		dp[i] = 1
		for j, b := range nums[:i] {
			if a > b {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		maxans = int(math.Max(float64(maxans), float64(dp[i])))
	}
	return maxans
}
