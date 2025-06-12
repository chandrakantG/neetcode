package main

import "fmt"

func main() {
	s := "zxyzxyz"
	s1 := "xxxa"
	fmt.Println(longestSubstringWithBruteForce(s))
	fmt.Println(longestSubstringWithMap(s))
	fmt.Println(longestSubstringWithBruteForce(s1))
	fmt.Println(longestSubstringWithMap(s1))
}

// time complexity: O(n*m)
// space complexity: O(m)
// Where n is the length of the string and m is the total number of unique characters in the string
func longestSubstringWithBruteForce(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		charSet := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if charSet[s[j]] {
				break
			}
			charSet[s[j]] = true
		}
		if len(charSet) > result {
			result = len(charSet)
		}
	}

	return result
}

// time complexity: O(n)
// space complexity: O(m)
// Where n is the length of the string and m is the total number of unique characters in the string
func longestSubstringWithMap(s string) int {
	result, l := 0, 0

	charSet := make(map[byte]bool)
	for r := 0; r < len(s); r++ {
		if charSet[s[r]] {
			delete(charSet, s[l])
			l++
		}
		charSet[s[r]] = true
	}
	result = len(charSet)

	return result
}
