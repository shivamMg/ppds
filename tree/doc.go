/*
Package tree provides implementation to get and print formatted tree.

Example:

	import "github.com/shivammg/ppds/tree"

	// a tree node.
	type Node struct {
		data int
		children []*Node
	}

	func (n *Node) Data() string {
		return strconv.Itoa(n.data)
	}

	// cannot return n.children directly.
	// https://github.com/golang/go/wiki/InterfaceSlice
	func (n *Node) Children() (c []tree.Node) {
		for _, child := range n.children {
			c = append(c, tree.Node(child))
		}
		return
	}

	// n1, n2 := Node{data: "b"}, Node{data: "c"}
	// n3 := Node{"a", []*Node{&n1, &n2}}
	// tree.Print(&n3) or tree.Sprint(&n3)
	// To pring horizontally:
	// tree.PrintHr(&n3) or tree.SprintHr(&n3)

*/
package tree
