package stack_test

import (
	"github.com/shivammg/ppds/stack"
	"testing"
)

type MyStack struct {
	elems []int
}

func (s *MyStack) Pop() (interface{}, bool) {
	l := len(s.elems)
	if l == 0 {
		return "", false
	}
	ele := s.elems[l-1]
	s.elems = s.elems[:l-1]
	return ele, true
}

func (s *MyStack) Push(ele interface{}) {
	s.elems = append(s.elems, ele.(int))
}

func TestPrint(t *testing.T) {
	s := MyStack{}
	s.Push(10)
	s.Push(19)
	s.Push(11)
	stack.Print(&s)
}
