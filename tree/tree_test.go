package tree_test

import (
	"testing"

	"github.com/shivammg/ppds/tree"
)

type Node struct {
	data string
	c    []*Node
}

func (n *Node) Data() string {
	return n.data
}

func (n *Node) Children() (children []tree.Node) {
	for _, c := range n.c {
		children = append(children, tree.Node(c))
	}
	return
}

func TestSprint(t *testing.T) {
	n1 := Node{data: "Femdom"}
	n2, n3, n16 := Node{"Dom&Sub", []*Node{&n1}}, Node{data: "Bondage"}, Node{data: "Masochism"}
	n4 := Node{"BDSM", []*Node{&n2, &n3, &n16}}
	n5, n6 := Node{data: "Age play"}, Node{data: "Furry"}
	n7 := Node{"Role enactment", []*Node{&n5, &n6}}
	n8 := Node{data: "Drugs"}
	n9 := Node{"Fetish", []*Node{&n7, &n8, &n4}}

	n12 := Node{data: "Vintage"}
	n11 := Node{"Classic", []*Node{&n12}}

	n13, n14 := Node{data: "Gonewild"}, Node{data: "Public"}
	n15 := Node{"Exhibition", []*Node{&n13, &n14}}
	n10 := Node{"Pr0n", []*Node{&n11, &n9, &n15}}

	const want = `Pr0n
├────────┬────────────────────────────────────────────────────┐         
Classic  Fetish                                               Exhibition
│        ├────────────────┬──────┐                            ├─────────┐     
Vintage  Role enactment   Drugs  BDSM                         Gonewild  Public
         ├─────────┐             ├────────┬────────┐        
         Age play  Furry         Dom&Sub  Bondage  Masochism
                                 │     
                                 Femdom
`

	got := tree.Sprint(&n10)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}
}

func TestSprintHor(t *testing.T) {
	n1, n2 := Node{data: "e"}, Node{data: "f"}
	n8 := Node{data: "h"}
	n3, n4 := Node{"d", []*Node{&n1, &n2}}, Node{"c", []*Node{&n8}}
	n5, n6 := Node{"b", []*Node{&n4, &n3}}, Node{data: "g"}
	n7 := Node{"a", []*Node{&n5, &n6}}

	want := `a┬b┬c─h
 │ └d┬e
 │   └f
 └g
`

	got := tree.SprintHr(&n7)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}
}
