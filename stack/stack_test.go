package stack_test

import (
	"github.com/shivammg/ppds/stack"
	"testing"
)

type MyStack struct {
	elems []int
}

func (s *MyStack) pop() (int, bool) {
	l := len(s.elems)
	if l == 0 {
		return 0, false
	}
	ele := s.elems[l-1]
	s.elems = s.elems[:l-1]
	return ele, true
}

func (s *MyStack) push(ele int) {
	s.elems = append(s.elems, ele)
}

func (s *MyStack) Pop() (interface{}, bool) {
	return s.pop()
}

func (s *MyStack) Push(ele interface{}) {
	s.push(ele.(int))
}

func TestPrint(t *testing.T) {
	s := MyStack{}
	s.push(10)
	s.push(199)
	s.push(1111)
	s.push(1111)
	s.push(12111)
	s.push(-113333111)
	s.push(1111)
	s.push(19)
	s.push(199)
	s.push(199)
	s.push(121499)
	s.push(-19019)
  want := `│    -19019│12│
│    121499│11│
│       199│10│
│       199│ 9│
│        19│ 8│
│      1111│ 7│
│-113333111│ 6│
│     12111│ 5│
│      1111│ 4│
│      1111│ 3│
│       199│ 2│
│        10│ 1│
└──────────┴──┘
`
  got := stack.Sprint(&s)
  if got != want {
    t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
  }
}
