package main

import (
	"fmt"
	"unicode"
)

func main() {
	ip := "Was it a car or a cat I saw"
	ip1 := "1rat tar1"
	fmt.Println(isPalindromeWithTwoPointer(ip))
	fmt.Println(isPalindromeWithTwoPointer(ip1))
}

// time complexity: O(n)
// space complexity: O(1)
func isPalindromeWithTwoPointer(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}
