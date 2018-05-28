package list

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

const (
	BoxVer       = "│"
	BoxHor       = "─"
	BoxDownLeft  = "┐"
	BoxDownRight = "┌"
	BoxUpLeft    = "┘"
	BoxUpRight   = "└"
	BoxVerLeft   = "┤"
	BoxVerRight  = "├"
)

// Node represents a linked list node.
type Node interface {
	// Data must return a string representing the value contained by the node.
	Data() string
	// Next must return the address of the next node. To mark the end of the list
	// return the zero value of the type.
	Next() Node
}

// Print prints the formatted linked list to standard output.
func Print(head Node) {
	fmt.Println(Sprint(head))
}

// Sprint returns the formatted linked list.
func Sprint(head Node) (s string) {
	data := []string{}
	for reflect.ValueOf(head) != reflect.Zero(reflect.TypeOf(head)) {
		data = append(data, head.Data())
		head = head.Next()
	}

	widths, maxDigitWidth := []int{}, digitWidth(len(data))
	for _, d := range data {
		w := utf8.RuneCountInString(d)
		if w > maxDigitWidth {
			widths = append(widths, w)
		} else {
			widths = append(widths, maxDigitWidth)
		}
	}

	for _, w := range widths {
		s += BoxDownRight + strings.Repeat(BoxHor, w) + BoxDownLeft
	}
	s += "\n"

	s += BoxVer
	for i, d := range data {
		s += fmt.Sprintf("%*s", widths[i], d)
		if i == len(data)-1 {
			s += BoxVer + "\n"
		} else {
			s += BoxVerRight + BoxVerLeft
		}
	}

	for _, w := range widths {
		s += BoxVerRight + strings.Repeat(BoxHor, w) + BoxVerLeft
	}
	s += "\n"

	s += BoxVer
	for i, w := range widths {
		s += fmt.Sprintf("%*d", w, i+1)
		if i == len(widths)-1 {
			s += BoxVer + "\n"
		} else {
			s += BoxVer + BoxVer
		}
	}

	for _, w := range widths {
		s += BoxUpRight + strings.Repeat(BoxHor, w) + BoxUpLeft
	}
	s += "\n"

	return
}

func digitWidth(d int) (w int) {
	for d != 0 {
		d = d / 10
		w++
	}
	return
}
