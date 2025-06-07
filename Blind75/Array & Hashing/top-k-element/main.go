package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	fmt.Println(topKElement(nums, k))
}

// 1. With Sort
// time complexity: O(n log n)
// space complexity: O(n)
func topKElement(nums []int, k int) []int {
	mNums := make(map[int]int)
	for _, v := range nums {
		mNums[v]++
	}
	fmt.Println("mNums:", mNums)

	arr := make([][2]int, 0, len(mNums))
	for num, cnt := range mNums {
		arr = append(arr, [2]int{cnt, num})
	}
	fmt.Println("arr:", arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	fmt.Println("arr:", arr)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}

	return res
}
