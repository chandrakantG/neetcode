package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagramWithMap("anagram", "nagaram"))
	fmt.Println(isAnagramWithSlice("anagram", "nagaram"))
}
func isAnagramWithSlice(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charSlice := make([]int, 26)

	for i := 0; i < len(s); i++ {
		charSlice[int(s[i]-'a')]++
		charSlice[int(t[i]-'a')]--
	}
	for _, v := range charSlice {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramWithMap(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[rune]int, len(s))

	for _, v := range s {
		charMap[v]++
	}

	for _, v := range t {
		if num, exist := charMap[v]; exist && num > 0 {
			charMap[v]--
		} else {
			return false
		}
	}

	return true
}
