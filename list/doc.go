/*
Package list provides implementation to get and print formatted linked list.

Example:

	import "github.com/shivammg/ppds/list"

	// a linked list node. implements list.Node.
	type Node struct {
		data int
		next *Node
	}

	func (n Node) Data() string {
		return strconv.Itoa(n.data)
	}

	func (n Node) Next() list.Node {
		return n.next
	}

	// n1 := Node{data: 11}
	// n2 := Node{12, &n1}
	// n3 := Node{13, &n2}
	// list.Print(n3) or list.Sprint(n3)

*/
package list
