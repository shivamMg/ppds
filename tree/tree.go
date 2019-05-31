package tree

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	BoxVer       = "│"
	BoxHor       = "─"
	BoxVerRight  = "├"
	BoxDownLeft  = "┐"
	BoxDownRight = "┌"
	BoxDownHor   = "┬"
	BoxUpRight   = "└"
	// Gutter is number of spaces between two adjacent child nodes
	Gutter = 2
)

// Node represents a node in a tree.
type Node interface {
	// Data must return a value representing the node.
	Data() interface{}
	// Children must return a list of all child nodes of the node.
	Children() []Node
}

type queue struct {
	arr []Node
}

func (q queue) empty() bool {
	return len(q.arr) == 0
}

func (q queue) len() int {
	return len(q.arr)
}

func (q *queue) push(n Node) {
	q.arr = append(q.arr, n)
}

func (q *queue) pop() Node {
	if q.empty() {
		return nil
	}
	ele := q.arr[0]
	q.arr = q.arr[1:]
	return ele
}

func (q *queue) peek() Node {
	if q.empty() {
		return nil
	}
	return q.arr[0]
}

// Print prints the formatted tree to standard output.
func Print(root Node) {
	fmt.Print(Sprint(root))
}

// Sprint returns the formatted tree.
func Sprint(root Node) string {
	parents := map[Node]Node{}
	setParents(parents, root)
	isLeftMostChild := func(n Node) bool {
		p, ok := parents[n]
		if !ok {
			// root
			return true
		}
		return p.Children()[0] == n
	}

	paddings := map[Node]int{}
	setPaddings(paddings, map[Node]int{}, 0, root)

	q := queue{}
	q.push(root)
	lines := []string{}
	for !q.empty() {
		// line storing branches, and line storing nodes
		branches, nodes := "", ""
		// runes covered
		covered := 0
		qLen := q.len()
		for i := 0; i < qLen; i++ {
			n := q.pop()
			for _, c := range n.Children() {
				q.push(c)
			}

			spaces := paddings[n] - covered
			// Should always have some spacing
			if spaces < 0 {
				spaces = 1
			}
			data := fmt.Sprintf("%v", n.Data())
			nodes += strings.Repeat(" ", spaces) + data

			w := utf8.RuneCountInString(data)
			covered += spaces + w
			current, next := isLeftMostChild(n), isLeftMostChild(q.peek())
			if current {
				branches += strings.Repeat(" ", spaces)
			} else {
				branches += strings.Repeat(BoxHor, spaces)
			}

			if current && next {
				branches += BoxVer
			} else if current {
				branches += BoxVerRight
			} else if next {
				branches += BoxDownLeft
			} else {
				branches += BoxDownHor
			}

			if next {
				branches += strings.Repeat(" ", w-1)
			} else {
				branches += strings.Repeat(BoxHor, w-1)
			}
		}
		lines = append(lines, branches, nodes)
	}

	s := ""
	// ignore first line since it's the branch above root
	for _, line := range lines[1:] {
		s += strings.TrimRight(line, " ") + "\n"

	}
	return s
}

// setPaddings sets left padding (distance of a node from the root)
// for each node in the tree.
func setPaddings(paddings map[Node]int, widths map[Node]int, pad int, root Node) {
	for _, c := range root.Children() {
		paddings[c] = pad
		setPaddings(paddings, widths, pad, c)
		pad += width(widths, c)
	}
}

// setParents sets child-parent relationships for the tree rooted
// at root.
func setParents(parents map[Node]Node, root Node) {
	for _, c := range root.Children() {
		parents[c] = root
		setParents(parents, c)
	}
}

// width returns either the sum of widths of it's children or its own
// data length depending on which one is bigger. widths is used in
// memoization.
func width(widths map[Node]int, n Node) int {
	if w, ok := widths[n]; ok {
		return w
	}

	w := utf8.RuneCountInString(fmt.Sprintf("%v", n.Data())) + Gutter
	widths[n] = w
	if len(n.Children()) == 0 {
		return w
	}

	sum := 0
	for _, c := range n.Children() {
		sum += width(widths, c)
	}
	if sum > w {
		widths[n] = sum
		return sum
	}
	return w
}
