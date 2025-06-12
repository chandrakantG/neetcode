package main

import "fmt"

func main() {
	prices := []int{10, 1, 5, 6, 7, 1}
	prices1 := []int{10, 8, 7, 5, 2}
	fmt.Println(buyAndSellWithBruteForce(prices))
	fmt.Println(buyAndSellWithTwoPointer(prices))
	fmt.Println(buyAndSellWithBruteForce(prices1))
	fmt.Println(buyAndSellWithTwoPointer(prices1))
}

// time complexity : O(n2)
// space complexity: O(1)
func buyAndSellWithBruteForce(prices []int) int {
	result := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices)-1; j++ {
			profit := prices[j] - prices[i]
			if profit > result {
				result = profit
			}
		}
	}

	return result
}

// time complexity : O(n)
// space complexity: O(1)
func buyAndSellWithTwoPointer(prices []int) int {
	result := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > result {
				result = profit
			}
		} else {
			l = r
		}
		r++
	}
	return result
}
