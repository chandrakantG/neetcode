package main

import "fmt"

func main() {
	n := 5
	n1 := 6
	fmt.Println(climbStairsWithBottomUp(n))
	fmt.Println(climbStairsWithBottomUp(n1))
	fmt.Println(climbStairs(n))
	fmt.Println(climbStairs(n1))
}

// time complexity: O(n)
// space complexity: O(n)
func climbStairsWithBottomUp(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	fmt.Println(dp)
	return dp[n-1]
}

// time complexity: O(n)
// space complexity: O(1)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	one, two := 1, 1
	for i := 1; i < n; i++ {
		one, two = one+two, one
	}
	return one
}
