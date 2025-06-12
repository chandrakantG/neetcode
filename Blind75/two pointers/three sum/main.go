package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumWithBruteForce(nums))
	fmt.Println(threeSumWithTwoPointer(nums))
}

// time complexity: O(n3)
// space complexity: O(m)- m number of triplates
func threeSumWithBruteForce(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	n := len(nums)
	resp := make(map[[3]int]struct{})
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					resp[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}
	result := [][]int{}
	for sum := range resp {
		result = append(result, []int{sum[0], sum[1], sum[2]})
	}
	return result
}

// time complexity: O(n2)
// space complexity: O(m)- m number of triplates
func threeSumWithTwoPointer(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	fmt.Println("nums:", nums)
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		fmt.Println("a:", a)
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := a + nums[l] + nums[r]
			fmt.Println("threeSum:", threeSum)
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}
