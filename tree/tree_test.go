package tree_test

import (
	"testing"

	"github.com/shivamMg/ppds/tree"
)

type Node struct {
	data string
	c    []*Node
}

func (n Node) Data() interface{} {
	return n.data
}

func (n Node) Children() (children []tree.Node) {
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

func TestSprintAndSprintHr(t *testing.T) {
	var classical, popRock Node
	{
		n1, n2 := Node{data: "Light"}, Node{data: "Heavy"}
		n3, n4 := Node{data: "Piano"}, Node{"Orchestra", []*Node{&n1, &n2}}
		n5, n6 := Node{data: "Male"}, Node{data: "Female"}
		n7, n8 := Node{"Opera", []*Node{&n5, &n6}}, Node{data: "Chorus"}
		n9, n10 := Node{"Instrumental", []*Node{&n3, &n4}}, Node{"Vocal", []*Node{&n7, &n8}}
		classical = Node{"Classical", []*Node{&n9, &n10}}
	}
	{
		n3 := Node{data: "Heavy metal"}
		n4, n5 := Node{data: "Dancing"}, Node{data: "Soft"}
		n6, n7 := Node{"Rock", []*Node{&n3}}, Node{"Country", []*Node{&n4, &n5}}
		n8, n9 := Node{data: "Late pop"}, Node{data: "Disco"}
		n10, n11 := Node{data: "Soft techno"}, Node{data: "Hard techno"}
		n12, n13 := Node{"Pop", []*Node{&n8, &n9}}, Node{"Techno", []*Node{&n10, &n11}}
		n14, n15 := Node{"Organic", []*Node{&n6, &n7}}, Node{"Electronic", []*Node{&n12, &n13}}
		popRock = Node{"Pop/Rock", []*Node{&n14, &n15}}
	}
	music := Node{"Music", []*Node{&classical, &popRock}}

	want := `Music
├──────────────────────────────────────────┐
Classical                                  Pop/Rock
├────────────────────┐                     ├───────────────────────────┐
Instrumental         Vocal                 Organic                     Electronic
├──────┐             ├─────────────┐       ├────────────┐              ├────────────────┐
Piano  Orchestra     Opera         Chorus  Rock         Country        Pop              Techno
       ├──────┐      ├─────┐               │            ├────────┐     ├─────────┐      ├────────────┐
       Light  Heavy  Male  Female          Heavy metal  Dancing  Soft  Late pop  Disco  Soft techno  Hard techno
`
	got := tree.Sprint(&music)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}

	want = `Music ┬─ Classical ┬─ Instrumental ┬─ Piano
      │            │               └─ Orchestra ┬─ Light
      │            │                            └─ Heavy
      │            └─ Vocal ┬─ Opera ┬─ Male
      │                     │        └─ Female
      │                     └─ Chorus
      └─ Pop/Rock ┬─ Organic ┬─ Rock ── Heavy metal
                  │          └─ Country ┬─ Dancing
                  │                     └─ Soft
                  └─ Electronic ┬─ Pop ┬─ Late pop
                                │      └─ Disco
                                └─ Techno ┬─ Soft techno
                                          └─ Hard techno
`
	got = tree.SprintHr(&music)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}

	want = `Music
├─ Classical
│  ├─ Instrumental
│  │  ├─ Piano
│  │  └─ Orchestra
│  │     ├─ Light
│  │     └─ Heavy
│  └─ Vocal
│     ├─ Opera
│     │  ├─ Male
│     │  └─ Female
│     └─ Chorus
└─ Pop/Rock
   ├─ Organic
   │  ├─ Rock
   │  │  └─ Heavy metal
   │  └─ Country
   │     ├─ Dancing
   │     └─ Soft
   └─ Electronic
      ├─ Pop
      │  ├─ Late pop
      │  └─ Disco
      └─ Techno
         ├─ Soft techno
         └─ Hard techno
`
	got = tree.SprintHrn(&music)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}
}

func TestSprintEmptyData(t *testing.T) {
	n2, n4, n5 := Node{data: "node2"}, Node{data: "node4"}, Node{data: "node5"}
	n3 := Node{data: "", c: []*Node{&n4}}
	n1 := Node{"node1", []*Node{&n2, &n3, &n5}}

	const want = `node1
├──────┬──────┐
node2         node5
       │
       node4
`

	got := tree.Sprint(&n1)
	if got != want {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", want, got)
	}
}
