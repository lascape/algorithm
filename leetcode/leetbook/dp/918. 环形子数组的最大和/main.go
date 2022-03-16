package main

import "fmt"

func main() {
	maxSubarraySumCircular([]int{5, -3, 5})
}

func maxSubarraySumCircular(nums []int) (ans int) {
	n := len(nums)
	ans = nums[0]
	cur := nums[0]
	for i := 1; i < n; i++ {
		cur = nums[i] + max(cur, 0)
		ans = max(ans, cur)
	}

	fmt.Println(ans)

	rightSums := make([]int, n)
	rightSums[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightSums[i] = rightSums[i+1] + nums[i]
	}
	fmt.Println(rightSums)

	maxRight := make([]int, n)
	maxRight[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], rightSums[i])
	}
	fmt.Println(maxRight)

	leftSum := 0
	for i := 0; i < n-2; i++ {
		leftSum += nums[i]
		ans = max(ans, leftSum+maxRight[i+2])
	}

	fmt.Println(ans)
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
