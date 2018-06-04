package stack_test

import (
	"testing"

	"github.com/shivammg/ppds/stack"
)

type myStack struct {
	elems []int
}

func (s *myStack) pop() (int, bool) {
	l := len(s.elems)
	if l == 0 {
		return 0, false
	}
	ele := s.elems[l-1]
	s.elems = s.elems[:l-1]
	return ele, true
}

func (s *myStack) push(ele int) {
	s.elems = append(s.elems, ele)
}

func (s *myStack) Pop() (interface{}, bool) {
	return s.pop()
}

func (s *myStack) Push(ele interface{}) {
	s.push(ele.(int))
}

func TestSprint(t *testing.T) {
	s := myStack{}
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
	want := `│     -19019 │12│
│     121499 │11│
│        199 │10│
│        199 │ 9│
│         19 │ 8│
│       1111 │ 7│
│ -113333111 │ 6│
│      12111 │ 5│
│       1111 │ 4│
│       1111 │ 3│
│        199 │ 2│
│         10 │ 1│
└────────────┴──┘
`
	got := stack.Sprint(&s)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}
}

type myStack2 struct {
	elems []string
}

func (s *myStack2) pop() (string, bool) {
	l := len(s.elems)
	if l == 0 {
		return "", false
	}
	ele := s.elems[l-1]
	s.elems = s.elems[:l-1]
	return ele, true
}

func (s *myStack2) push(ele string) {
	s.elems = append(s.elems, ele)
}

func (s *myStack2) Pop() (interface{}, bool) {
	return s.pop()
}

func (s *myStack2) Push(ele interface{}) {
	s.push(ele.(string))
}

func TestRunes(t *testing.T) {
	s := myStack2{}
	s.push("你好，世界")
	s.push("Hello world")
	s.push("Ciao mondo")
	s.push("नमस्ते दुनिया")
	want := `│ नमस्ते दुनिया │4│
│    Ciao mondo │3│
│   Hello world │2│
│         你好，世界 │1│
└───────────────┴─┘
`
	got := stack.Sprint(&s)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}
}
