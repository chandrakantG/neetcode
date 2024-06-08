package structures

type Stack struct {
	nums []int
}

// create new pointer stack
func NewStack() *Stack {
	return &Stack{
		nums: []int{},
	}
}

// append new element in stack
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// return last element from stack
func (s *Stack) Pop() int {
	sLen := len(s.nums)
	res := s.nums[sLen-1]
	s.nums = s.nums[:sLen-1]
	return res
}

// return len of stack
func (s *Stack) Len() int {
	return len(s.nums)
}

// check stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
