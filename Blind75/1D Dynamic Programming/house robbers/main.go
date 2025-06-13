package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 3}
	fmt.Println(rob(nums))
	nums1 := []int{2, 9, 8, 3, 6}
	fmt.Println(rob(nums1))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	fmt.Println(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	fmt.Println(dp)
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
