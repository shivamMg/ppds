package stack

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	BoxVer     = "│"
	BoxHor     = "─"
	BoxUpHor   = "┴"
	BoxUpLeft  = "┘"
	BoxUpRight = "└"
)

// Stack represents a stack of elements.
type Stack interface {
	// Pop must pop the top element out of the stack and return it. In case of
	// an empty stack ok should be false, else true.
	Pop() (ele interface{}, ok bool)
	// Push must insert an element in the stack. Since ele is of interface type,
	// type assertion must be done before inserting element in the stack.
	Push(ele interface{})
}

// Print prints the formatted stack to standard output.
func Print(s Stack) {
	fmt.Print(Sprint(s))
}

// Sprint returns the formatted stack.
func Sprint(s Stack) (str string) {
	elems := []interface{}{}
	e, ok := s.Pop()
	for ok {
		elems = append(elems, e)
		e, ok = s.Pop()
	}

	for i := len(elems) - 1; i >= 0; i-- {
		s.Push(elems[i])
	}

	data, maxWidth := []string{}, 0
	for _, e := range elems {
		d := fmt.Sprintf("%v", e)
		data = append(data, d)
		w := utf8.RuneCountInString(d)
		if w > maxWidth {
			maxWidth = w
		}
	}
	// position column maxWidth
	posW := digitWidth(len(data))

	for i, d := range data {
		str += fmt.Sprintf("%s%*s%s%*d%s\n", BoxVer, maxWidth, d, BoxVer, posW, len(data)-i, BoxVer)
	}
	str += fmt.Sprint(BoxUpRight, strings.Repeat("─", maxWidth), BoxUpHor, strings.Repeat("─", posW), BoxUpLeft, "\n")

	return
}

func digitWidth(d int) (w int) {
	for d != 0 {
		d = d / 10
		w++
	}
	return
}
