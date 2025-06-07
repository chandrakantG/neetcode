package main

import (
	"fmt"
	"sort"
)

func main() {
	ipStrings := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagramWithSort(ipStrings))

	fmt.Println(groupAnagramWithMap(ipStrings))
}

// 1. With Sort
// time complexity: O(m*nlogn) => m are length of string & n is length of longest string
// space complexity: O(m*n)
func groupAnagramWithSort(ipStrings []string) [][]string {
	sortedM := make(map[string][]string)
	for _, v := range ipStrings {
		sortedS := sortString(v)
		sortedM[sortedS] = append(sortedM[sortedS], v)
	}
	outPut := [][]string{}
	for _, group := range sortedM {
		outPut = append(outPut, group)
	}

	return outPut
}

func sortString(s string) string {
	characters := []rune(s)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}

// 2. With Map
// time complexity: O(m*nlogn) => m are length of string & n is length of longest string
// space complexity: O(m*n)
func groupAnagramWithMap(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		res[count] = append(res[count], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}
