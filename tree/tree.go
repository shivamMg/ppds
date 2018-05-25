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

	q := queue{}
	q.push(&n11)
	pop := 1
	for !q.empty() {
		nextPop := 0
		var prev *Node
		covered := 0
		for i := 0; i < pop; i++ {
			n := q.pop()
			spaces := 0
			if prev != nil {
				a, _, c2 := commonAncestor(parents, prev, n)
				// fmt.Println("ancestor", a.data, c1.data, c2.data)
				total := 0
				for _, c := range a.c {
					if c == c2 {
						break
					}
					total += width(c)
				}
				fmt.Println("covered", covered)
				spaces += total - covered
				// fmt.Println("spaces", n.data, spaces)
			}

			w := width(n)
			fmt.Printf("%s%*s", strings.Repeat(" ", spaces), w, n.data)
			covered += w

			for _, c := range n.c {
				nextPop++
				q.push(c)
			}
			prev = n
		}
		pop = nextPop
		fmt.Println()
	}
}

// common ancestor and it's two children that contain node1 and node2
func commonAncestor(parents map[*Node]*Node, node1, node2 *Node) (*Node, *Node, *Node) {
	n1, n2 := node1, node2
	// parent->child
	traversed1, traversed2 := map[*Node]*Node{}, map[*Node]*Node{}

	ok := true
	traversed1[n1] = nil
	for {
		traversed1[parents[n1]] = n1
		n1, ok = parents[n1]
		if !ok {
			break
		}
	}

	traversed2[n2] = nil
	for {
		traversed2[parents[n2]] = n2
		n2, ok = parents[n2]
		if !ok {
			break
		}
	}

	// traverse n1 until root or merge point
	n := node1
	for {
		c1, ok1 := traversed1[n]
		c2, ok2 := traversed2[n]
		if ok1 && ok2 {
			return n, c1, c2
		}
		n, ok = parents[n]
		if !ok {
			break
		}
	}
	return nil, nil, nil
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
