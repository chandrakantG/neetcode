package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 20, 4, 10, 3, 4, 5}
	nums2 := []int{0, 3, 2, 5, 4, 6, 1, 1}
	fmt.Println(longestConsecutiveWithBruteForce(nums))
	fmt.Println(longestConsecutiveWithBruteForce(nums2))
	fmt.Println(longestConsecutiveWithSort(nums))
	fmt.Println(longestConsecutiveWithSort(nums2))
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive(nums2))
}

// 1. With Brute Force
// time complexity: O(n2)
// space complexity: O(n)
func longestConsecutiveWithBruteForce(nums []int) int {
	var result int
	if len(nums) == 0 {
		return 0
	}

	store := make(map[int]struct{})
	for _, num := range nums {
		store[num] = struct{}{}
	}
	fmt.Println(store)

	for _, num := range nums {
		length, currNum := 0, num
		for _, ok := store[currNum]; ok; _, ok = store[currNum] {
			length++
			currNum++
		}
		// for {
		// 	if _, ok := store[currNum]; ok {
		// 		length++
		// 		currNum++
		// 		continue
		// 	}
		// 	break
		// }
		if length > result {
			result = length
		}
	}

	return result
}

// 2. With Sort
// time complexity: O(nlogn)
// space complexity: O(n)
func longestConsecutiveWithSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var result int
	sort.Ints(nums)
	fmt.Println(nums)

	curr, length, i := nums[0], 0, 0

	for i < len(nums) {
		if curr != nums[i] {
			curr = nums[i]
			length = 0
		}
		for i < len(nums) && nums[i] == curr {
			i++
		}
		curr++
		length++
		if length > result {
			result = length
		}
	}
	return result
}

func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	res := 0

	for _, num := range nums {
		if mp[num] == 0 {
			left := mp[num-1]
			fmt.Println("left:", left)
			right := mp[num+1]
			fmt.Println("right:", right)
			sum := left + right + 1
			mp[num] = sum
			mp[num-left] = sum
			mp[num+right] = sum
			fmt.Println("mp:", mp)
			fmt.Println("--------------")
			if sum > res {
				res = sum
			}
		}
	}
	return res
}
