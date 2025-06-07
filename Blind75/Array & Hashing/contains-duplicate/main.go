package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(hasDuplicateWithBruteForce(nums))
	fmt.Println(hasDuplicateWithSorting(nums))
	fmt.Println(hasDuplicateWithMap(nums))
}

// 1. With Brute Force
// time complexity: O(n2)
// space complexity: O(1)
func hasDuplicateWithBruteForce(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// 2. With Sorting
// time complexity: O(nlogn)
// space complexity: O(1) or O(n) depending on the sorting algorithm.
func hasDuplicateWithSorting(nums []int) bool {
	sort.Ints(nums)
	fmt.Println("sorted:", nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// 3. With Map
// time complexity: O(n)
// space complexity: O(n)
func hasDuplicateWithMap(nums []int) bool {
	existNum := make(map[int]bool)
	for _, v := range nums {
		if _, ok := existNum[v]; ok {
			return true
		}
		existNum[v] = true
	}
	return false
}
