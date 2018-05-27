package main

import (
	"fmt"
	"strings"
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

	leftMostBranch := getLeftMostBranch(&n11)

	q := queue{}
	q.push(&n11)
	qLen := 1
	for !q.empty() {
		var prevNode *Node
		pushed := 0
		for i := 0; i < qLen; i++ {
			n := q.pop()
			for _, c := range n.c {
				pushed++
				q.push(c)
			}

			if isLeftMostChild(parents, n) {
				if i == 0 {
					fmt.Print(n.data)
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
				// traverse through prevNode's ancestory to get a
				// node in the left-most branch of the tree
				node := prevNode
				for {
					if _, ok := leftMostBranch[node]; ok {
						break
					}
					node = parents[node]
				}
				covered := width(node)
				spaces := leftPad - covered
				fmt.Print(strings.Repeat(" ", spaces), n.data)
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
			spaces := width(prev) - 1
			fmt.Print(strings.Repeat(" ", spaces), n.data)
			prevNode = n
		}
		qLen = pushed
		fmt.Println()
	}
}

func getLeftMostBranch(root *Node) map[*Node]struct{} {
	b := make(map[*Node]struct{})
	b[root] = struct{}{}
	for len(root.c) != 0 {
		root = root.c[0]
		b[root] = struct{}{}
	}
	return b
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
	if len(n.c) == 0 {
		return 1
	}
	sum := 0
	for _, c := range n.c {
		sum += width(c)
	}
	return sum
}
