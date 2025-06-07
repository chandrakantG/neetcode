package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 6}
	nums1 := []int{-1, 0, 1, 2, 3}
	fmt.Println(productExceptSelfWithBruteForce(nums))
	fmt.Println(productExceptSelfWithBruteForce(nums1))

	fmt.Println(productExceptSelfWithDivision(nums))
	fmt.Println(productExceptSelfWithDivision(nums1))
}

// 1. With Brute Force
// time complexity: O(n2)
// space complexity: O(n)
func productExceptSelfWithBruteForce(nums []int) []int {
	n := len(nums)

	result := make([]int, n)
	fmt.Printf("%v,%v,%v \n", result, cap(result), len(result))

	// result1 := []int{}
	// fmt.Printf("%v,%v,%v \n", result1, cap(result1), len(result1))
	for i := 0; i < n; i++ {
		prod := 1
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			prod *= nums[j]
		}
		result[i] = prod
		// result1 = append(result1, prod)
	}
	return result
}

// 2. With Division
// time complexity: O(n)
// space complexity: O(n)
func productExceptSelfWithDivision(nums []int) []int {
	prod := 1
	zeroCount := 0

	for _, num := range nums {
		if num != 0 {
			prod *= num
		} else {
			zeroCount++
		}
	}
	fmt.Println("prod:", prod)
	fmt.Println("zeroCount:", zeroCount)
	res := make([]int, len(nums))
	if zeroCount > 1 {
		return res
	}
	fmt.Println("res:", res)
	for i, num := range nums {
		if zeroCount > 0 {
			if num == 0 {
				res[i] = prod
			} else {
				res[i] = 0 // have 0 element so product is 0
			}
		} else {
			res[i] = prod / num
		}
	}
	return res
}
