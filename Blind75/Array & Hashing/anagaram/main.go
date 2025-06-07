package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := "jar"
	s2 := "jam"
	s3 := "racecar"
	s4 := "carrace"
	fmt.Println(isAnagramWithSort(s1, s2))
	fmt.Println(isAnagramWithSort(s3, s4))

	fmt.Println(isAnagramWithMap(s1, s2))
	fmt.Println(isAnagramWithMap(s3, s4))

	fmt.Println(isAnagramWithArray(s1, s2))
	fmt.Println(isAnagramWithArray(s3, s4))
}

// 1. With Sort
// time complexity: O(nlogn +mlogm) => n & m are lenght of strings
// space complexity: O(n+m)
func isAnagramWithSort(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Runes, s2Runes := []rune(s1), []rune(s2)

	sort.Slice(s1Runes, func(i, j int) bool {
		return s1Runes[i] < s1Runes[j]
	})

	sort.Slice(s2Runes, func(i, j int) bool {
		return s2Runes[i] < s2Runes[j]
	})

	for i := range len(s1Runes) {
		if s1Runes[i] != s2Runes[i] {
			return false
		}
	}
	return true
}

// 2. With Map
// time complexity: O(n+m) => n & m are lenght of strings
// space complexity: O(1) since we have at most 26 different characters.
func isAnagramWithMap(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Map, s2Map := make(map[rune]int), make(map[rune]int)

	for i, ch := range s1 {
		s1Map[ch]++
		s2Map[rune(s2[i])]++
	}
	fmt.Println(s1Map)
	fmt.Println(s2Map)

	for k, v := range s1Map {
		if s2Map[k] != v {
			return false
		}
	}

	return true
}

// 2. With Array
// time complexity: O(n+m) => n & m are lenght of strings
// space complexity: O(1) since we have at most 26 different characters.
func isAnagramWithArray(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := [26]int{}
	for k, v := range s1 {
		count[v-'a']++
		count[s2[k]-'a']--
	}

	fmt.Println(count)

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
