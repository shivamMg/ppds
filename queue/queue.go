package queue

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	Arrow        = "→"
	BoxVer       = "│"
	BoxHor       = "─"
	BoxDownLeft  = "┐"
	BoxDownRight = "┌"
	BoxUpLeft    = "┘"
	BoxUpRight   = "└"
	BoxVerLeft   = "┤"
	BoxVerRight  = "├"
)

// Queue represents a queue of elements.
type Queue interface {
	// Dequeue must return the first (FIFO) element out of the queue. If queue is
	// empty ok should be false, else true.
	Dequeue() (ele interface{}, ok bool)
	// Enqueue must insert ele in the queue. Since ele is of interface type, type
	// assertion must be done before inserting inside the queue.
	Enqueue(ele interface{})
}

// Print prints the formatted queue to standard output.
func Print(q Queue) {
	fmt.Print(Sprint(q))
}

// Sprint returns the formatted queue.
func Sprint(q Queue) (s string) {
	elems := []interface{}{}
	e, ok := q.Dequeue()
	for ok {
		elems = append(elems, e)
		e, ok = q.Dequeue()
	}

	for _, e := range elems {
		q.Enqueue(e)
	}

	maxDigitWidth := digitWidth(len(elems))
	data, widths := []string{}, []int{}
	for _, e := range elems {
		d := fmt.Sprintf("%v", e)
		data = append(data, d)
		w := utf8.RuneCountInString(d)
		if w > maxDigitWidth {
			widths = append(widths, w)
		} else {
			widths = append(widths, maxDigitWidth)
		}
	}

	for _, w := range widths {
		s += " " + BoxDownRight + strings.Repeat(BoxHor, w) + BoxDownLeft
	}
	s += "\n"

	s += Arrow + BoxVer
	for i, d := range data {
		s += fmt.Sprintf("%*s", widths[i], d)
		if i == len(data)-1 {
			s += BoxVer + Arrow + "\n"
		} else {
			s += BoxVer + Arrow + BoxVer
		}
	}

	for _, w := range widths {
		s += " " + BoxVerRight + strings.Repeat(BoxHor, w) + BoxVerLeft
	}
	s += "\n"

	s += " " + BoxVer
	for i, w := range widths {
		s += fmt.Sprintf("%*d", w, i+1)
		if i == len(widths)-1 {
			s += BoxVer + "\n"
		} else {
			s += BoxVer + " " + BoxVer
		}
	}

	for _, w := range widths {
		s += " " + BoxUpRight + strings.Repeat(BoxHor, w) + BoxUpLeft
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
