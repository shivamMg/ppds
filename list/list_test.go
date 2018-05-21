package list_test

import (
  "testing"
  "strconv"
	"github.com/shivammg/ppds/list"
)

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

func TestPrint(t *testing.T) {
  n1 := Node{10, nil}
  n2 := Node{11, &n1}
  n3 := Node{13, &n2}
  list.Print(&n3)
}
