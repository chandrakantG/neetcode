package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 5, 6, 9}
	target := 9
	fmt.Println(combinationSum(nums, target))
}

func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	var dfs func(int, []int, int)

	dfs = func(i int, cur []int, total int) {
		if target == total {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		// if i >= len(nums) || total > target {
		// 	return
		// }
		// cur = append(cur, nums[i])
		// dfs(i, cur, total+nums[i])
		// cur = cur[:len(cur)-1]
		// dfs(i+1, cur, total)

		for j := i; j < len(nums); j++ {
			if total+nums[j] > target {
				return
			}
			cur = append(cur, nums[j])
			dfs(j, cur, total+nums[j])
			cur = cur[:len(cur)-1]
		}

	}

	dfs(0, []int{}, 0)
	return res
}
