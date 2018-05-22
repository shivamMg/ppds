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
  n1 := Node{935, nil}
  n2 := Node{-10324, &n1}
  n3 := Node{3000000, &n2}
  want := `┌───────┐┌──────┐┌───┐
│3000000├┤-10324├┤935│
├───────┤├──────┤├───┤
│      1││     2││  3│
└───────┘└──────┘└───┘
`
  got := list.Sprint(&n3)
  if want != got {
    t.Errorf("Expected: %s\nGot: %s\n", want, got)
  }
}
