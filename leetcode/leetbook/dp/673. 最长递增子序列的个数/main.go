package main

func findNumberOfLIS(nums []int) (ans int) {

	maxLen := 0
	n := len(nums)
	dp := make([]int, n)
	cnt := make([]int, n)

	for i, a := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, b := range nums[:i] {
			if a > b {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return
}
