package main

import "fmt"

func main() {
	s := "([{}])"
	s1 := "[(])"
	fmt.Println(isValid(s))
	fmt.Println(isValid(s1))
}

// time complexity : O(n)
// space complexity: O(n)
func isValid(s string) bool {
	closeMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	cStack := []rune{}
	for _, c := range s {
		if val, ok := closeMap[c]; ok {
			if len(cStack) > 0 && cStack[len(cStack)-1] == val {
				cStack = cStack[:len(cStack)-1]
			} else {
				return false
			}
		} else {
			cStack = append(cStack, c)
		}
	}
	return len(cStack) == 0
}
