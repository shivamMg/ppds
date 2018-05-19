package stack

import (
	"fmt"
)

// Stack represents a stack of elements.
type Stack interface {
	Pop() (interface{}, bool)
	Push(ele interface{})
}

func Print(s Stack) {
	elems := []interface{}{}
	for {
		if ele, ok := s.Pop(); ok {
			elems = append(elems, ele)
			continue
		}
		break
	}

	for _, ele := range elems {
		fmt.Println(ele)
	}

	for i := len(elems) - 1; i >= 0; i-- {
		s.Push(elems[i])
	}
}
