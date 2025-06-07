package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 6}
	target := 7
	fmt.Println(twoSumWithMap(nums, target))
	fmt.Println(twoSumWithBruteForce(nums, target))
}

// 1. With Map
// time complexity: O(n)
// space complexity: O(n)
func twoSumWithMap(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		diff := target - v
		if ind, ok := numMap[diff]; ok {
			return []int{ind, k}
		}
		numMap[v] = k
	}
	return []int{}
}

// 2. With Brute Force
// time complexity: O(n2)
// space complexity: O(1)
func twoSumWithBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}

		}
	}
	return []int{}
}
