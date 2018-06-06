/*
Package stack provides implementation to get and print formatted stack.

Example:

	import "github.com/shivamMg/ppds/stack"

	// a stack implementation.
	type myStack struct {
		elems []int
	}

	func (s *myStack) push(ele int) {
		s.elems = append(s.elems, ele)
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

	// myStack implements stack.Stack. Notice that the receiver is of *myStack
	// type - since Push and Pop are required to modify s.
	func (s *myStack) Pop() (interface{}, bool) {
		return s.pop()
	}

	func (s *myStack) Push(ele interface{}) {
		s.push(ele.(int))
	}

	// s := myStack{}
	// s.push(11)
	// s.push(12)
	// stack.Print(&s)

*/
package stack
