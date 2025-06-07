package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func main() {
	s := Solution{}
	strs1 := []string{"neet", "code", "love", "you"}
	encode := s.Encode(strs1)
	fmt.Println("encode:", encode)
	fmt.Println(s.Decode(encode))
	fmt.Println("-----------------")
	strs2 := []string{"we", "say", ":", "yes"}
	encode = s.Encode(strs2)
	fmt.Println("encode:", encode)
	fmt.Println(s.Decode(encode))
}

// 1. With Brute Force
// time complexity: O(m) , m: sum of length all strings, n:no of strings
// space complexity: O(m+n)

func (s Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sizes := []string{}
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}
	return strings.Join(sizes, ",") + "#" + strings.Join(strs, "")
}

func (s Solution) Decode(encode string) []string {
	if encode == "" {
		return []string{}
	}
	parts := strings.SplitN(encode, "#", 2)
	sizes := strings.Split(parts[0], ",")

	fmt.Println(parts)
	fmt.Println(sizes)

	var res []string
	i := 0
	for _, v := range sizes {
		length, _ := strconv.Atoi(v)
		min := i
		max := i + length
		str := parts[1][min:max]
		res = append(res, str)
		i = max
	}

	return res
}
