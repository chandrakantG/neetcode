package main

import "fmt"

func main() {
	s := "XYYX"
	k := 2
	fmt.Println(characterReplacement(s, k))
}

func characterReplacement(s string, k int) int {
	resp := 0
	countMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		countMap[s[i]]++
	}
	for c := range countMap {
		count, l := 0, 0
		for r := 0; r < len(s); r++ {
			if s[r] == c {
				count++
			}
			for (r-l+1)-count > k {
				if s[l] == c {
					count--
				}
				l++
			}
			resp = max(resp, r-l+1)
		}
	}

	return resp
}
