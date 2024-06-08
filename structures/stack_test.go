package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	ast := assert.New(t)

	s := NewStack()
	ast.True(s.IsEmpty(), "check new stack s is empty")

	start, end := 0, 100
	for i := start; i < end; i++ {
		s.Push(i)
		ast.Equal(s.Len(), i+1, "check length of s after push")
	}
	for i := end; i > 0; i-- {
		ast.Equal(i-1, s.Pop(), "pop number from s")
	}
	ast.True(s.IsEmpty(), "check stack is empty after pops")
}
