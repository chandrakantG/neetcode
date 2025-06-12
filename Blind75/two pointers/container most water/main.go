package main

import "fmt"

func main() {
	height := []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Println(maxAreaWithBruteForce(height))
	fmt.Println(maxAreaWithTwoPointer(height))
}

// time complexity: O(n2)
// space complexity: O(1)
func maxAreaWithBruteForce(heights []int) int {
	res := 0
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			area := min(heights[i], heights[j]) * (j - i)
			if area > res {
				res = area
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// time complexity: O(n)
// space complexity: O(1)
func maxAreaWithTwoPointer(heights []int) int {
	result := 0
	l, r := 0, len(heights)-1

	for l < r {
		area := min(heights[l], heights[r]) * (r - l)
		if area > result {
			result = area
		}
		if heights[l] <= heights[r] {
			l++
		} else {
			r--
		}
	}

	return result
}
