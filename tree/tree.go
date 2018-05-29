package main

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
)

type Node struct {
	data string
	c    []*Node
}

type queue struct {
	arr []*Node
}

func (q queue) empty() bool {
	return len(q.arr) == 0
}

func (q *queue) push(n *Node) {
	q.arr = append(q.arr, n)
}

func (q *queue) pop() *Node {
	if q.empty() {
		return nil
	}
	ele := q.arr[0]
	q.arr = q.arr[1:]
	return ele
}

func (q *queue) peek() *Node {
	if q.empty() {
		return nil
	}
	return q.arr[0]
}

func main() {
	n1, n2 := Node{data: "1"}, Node{data: "2"}
	n3 := Node{"5", []*Node{&n1, &n2}}
	n4, n5 := Node{data: "8"}, Node{data: "6"}
	n6 := Node{"3", []*Node{&n3, &n4, &n5}}

	n7, n8 := Node{data: "8"}, Node{data: "9"}
	n9 := Node{"7", []*Node{&n7, &n8}}
	n10 := Node{"4", []*Node{&n9}}

	n11 := Node{"2", []*Node{&n6, &n10}}

	// get parents
	parents := map[*Node]*Node{}
	getParents(parents, &n11)

	q := queue{}
	q.push(&n11)
	qLen := 1
	isRoot := true
	for !q.empty() {
		var prevNode *Node
		pushed := 0
		covered := 0
		branchLine := ""
		line := ""
		for i := 0; i < qLen; i++ {
			n := q.pop()
			for _, c := range n.c {
				pushed++
				q.push(c)
			}

			if isLeftMostChild(parents, n) {
				if i == 0 {
					// fmt.Print(n.data)
					line += n.data
					covered = utf8.RuneCountInString(n.data)
					if !isRoot {
						if isLeftMostChild(parents, q.peek()) {
							branchLine += BoxVer
						} else {
							branchLine += BoxVerRight
						}
						branchLine += strings.Repeat(BoxHor, covered-1)
					}
					isRoot = false
					prevNode = n
					continue
				}
				anc, _, c2 := commonAncestor(parents, prevNode, n)
				leftPad := 0
				for _, c := range anc.c {
					if c == c2 {
						break
					}
					leftPad += width(c)
				}
				spaces := leftPad - covered
				// fmt.Print(strings.Repeat(" ", spaces), n.data)
				line += strings.Repeat(" ", spaces) + n.data
				w := utf8.RuneCountInString(n.data)
				covered = spaces + w
				branchLine += strings.Repeat(" ", spaces)
				if isLeftMostChild(parents, q.peek()) {
					branchLine += BoxVer
				} else {
					branchLine += BoxVerRight
				}
				branchLine += strings.Repeat(BoxHor, w-1)
				prevNode = n
				continue
			}
			p := parents[n]
			prev := p.c[0]
			for _, c := range p.c[1:] {
				if c == n {
					break
				}
				prev = c
			}
			spaces := width(prev) - utf8.RuneCountInString(prev.data)
			// fmt.Print(strings.Repeat(" ", spaces), n.data)
			line += strings.Repeat(" ", spaces) + n.data
			w := utf8.RuneCountInString(n.data)
			covered += spaces + w
			branchLine += strings.Repeat(BoxHor, spaces)
			if isLeftMostChild(parents, q.peek()) {
				branchLine += BoxDownLeft
			} else {
				branchLine += BoxDownHor
			}
			branchLine += strings.Repeat(BoxHor, w-1)
			prevNode = n
		}
		qLen = pushed
		fmt.Println(branchLine)
		fmt.Println(line)
	}
}

func isLeftMostChild(parents map[*Node]*Node, node *Node) bool {
	p, ok := parents[node]
	if !ok {
		return true
	}
	return p.c[0] == node
}

// common ancestor and it's two children that contain node1 and node2
func commonAncestor(parents map[*Node]*Node, node1, node2 *Node) (*Node, *Node, *Node) {
	n1, n2 := node1, node2
	// parent->child
	traversed1, traversed2 := map[*Node]*Node{}, map[*Node]*Node{}

	traversed1[n1] = nil
	p, ok := parents[n1]
	for ok {
		traversed1[p] = n1
		n1 = p
		p, ok = parents[n1]
	}

	traversed2[n2] = nil
	p, ok = parents[n2]
	for ok {
		traversed2[p] = n2
		n2 = p
		p, ok = parents[n2]
	}

	// traverse upward until merge point
	n := node1
	for {
		_, ok1 := traversed1[n]
		_, ok2 := traversed2[n]
		if ok1 && ok2 {
			break
		}
		n = parents[n]
	}
	return n, traversed1[n], traversed2[n]
}

func getParents(parents map[*Node]*Node, n *Node) {
	for _, c := range n.c {
		parents[c] = n
		getParents(parents, c)
	}
}

func width(n *Node) int {
	w := utf8.RuneCountInString(n.data) + 1
	if len(n.c) == 0 {
		return w
	}
	sum := 0
	for _, c := range n.c {
		sum += width(c)
	}
	if sum > w {
		return sum
	}
	return w
}
