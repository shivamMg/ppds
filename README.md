# ppds [![godoc](https://godoc.org/github.com/shivammg/ppds?status.svg)](https://godoc.org/github.com/shivamMg/ppds) ![Build](https://github.com/shivamMg/trie/actions/workflows/ci.yml/badge.svg?branch=master) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Pretty Print Data Structures

Stacks, queues, trees and linked lists are data structures that you might find yourself working with quite often. This library lets you pretty print these with minimum effort. Certain assumptions can be made for each data structure, for instance, a stack or a queue will have methods synonymous to Push (insert) and Pop (remove), a tree node will have links to its child nodes, or a linked list node will have a link to its next adjacent node. This library utilises those assumptions, and exports interfaces and functions to pretty print them.

The following data structures are supported:

### Trees
[![tree](https://godoc.org/github.com/shivammg/ppds?status.svg)](https://godoc.org/github.com/shivamMg/ppds/tree)

A type that satisfies the following interface can be printed using `tree.Print`.

```
type Node interface {
	Data() interface{}
	Children() []Node
}
```

```
Music
├──────────────────────────────────────────┐
Classical                                  Pop/Rock
├────────────────────┐                     ├───────────────────────────┐
Instrumental         Vocal                 Organic                     Electronic
├──────┐             ├─────────────┐       ├────────────┐              ├────────────────┐
Piano  Orchestra     Opera         Chorus  Rock         Country        Pop              Techno
       ├──────┐      ├─────┐               │            ├────────┐     ├─────────┐      ├────────────┐
       Light  Heavy  Male  Female          Heavy metal  Dancing  Soft  Late pop  Disco  Soft techno  Hard techno
```

If branching factor is high or data lengths are large, it would make more sense to print the tree horizontally. That can be done using `tree.PrintHr` function.

```
Music ┬─ Classical ┬─ Instrumental ┬─ Piano
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
```

Inspired by [tree](http://mama.indstate.edu/users/ice/tree/), a popular directory listing command, `tree.PrintHrn` prints the tree in similar style. Every node is printed on a separate line making the output tree look less cluttered and more elegant. It does take up vertical space though.

```
Music
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
```

### Linked Lists
[![list](https://godoc.org/github.com/shivammg/ppds?status.svg)](https://godoc.org/github.com/shivamMg/ppds/list)

```
type Node interface {
	Data() interface{}
	Next() Node
}
```

A counter is also printed below each node so one doesn't have to manually count them by pointing a finger at the screen. Also, a list can contain a million elements, in which case counting by simply shifting finger on screen might not even be possible.

```
┌───┐┌───┐┌───┐┌─┐┌────┐┌────┐┌────┐
│333├┤222├┤111├┤0├┤-111├┤-222├┤-333│
├───┤├───┤├───┤├─┤├────┤├────┤├────┤
│  1││  2││  3││4││   5││   6││   7│
└───┘└───┘└───┘└─┘└────┘└────┘└────┘
```

### Stacks
[![stack](https://godoc.org/github.com/shivammg/ppds?status.svg)](https://godoc.org/github.com/shivamMg/ppds/stack)

```
type Stack interface {
	Pop() (ele interface{}, ok bool)
	Push(ele interface{})
}
```

```
│ factorial (1) │7│
│ factorial (2) │6│
│ factorial (3) │5│
│ factorial (4) │4│
│ factorial (5) │3│
│ factorial (6) │2│
│ factorial (7) │1│
└───────────────┴─┘
```

### Queues
[![queue](https://godoc.org/github.com/shivammg/ppds?status.svg)](https://godoc.org/github.com/shivamMg/ppds/queue)

```
type Queue interface {
	Pop() (ele interface{}, ok bool)
	Push(ele interface{})
}
```

```
 ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐
→│Dec│→│Nov│→│Oct│→│Sep│→│Aug│→│Jul│→│Jun│→│May│→│Apr│→│Mar│→│Feb│→│Jan│→
 ├───┤ ├───┤ ├───┤ ├───┤ ├───┤ ├───┤ ├───┤ ├───┤ ├───┤ ├───┤ ├───┤ ├───┤
 │  1│ │  2│ │  3│ │  4│ │  5│ │  6│ │  7│ │  8│ │  9│ │ 10│ │ 11│ │ 12│
 └───┘ └───┘ └───┘ └───┘ └───┘ └───┘ └───┘ └───┘ └───┘ └───┘ └───┘ └───┘
```

## Notes

Printing Stacks and Queues requires defining Pop and Push methods. Internally these methods are used to get all elements - first by popping all elements then pushing them back in correct order. Use it with caution since the actual data structure is being modified when printing. After printing it'll be back to its original state. I've discussed why I decided to keep it this way in this [reddit thread](https://redd.it/8pbvrd). If there is no simultaneous read/write while printing, for instance from a different go routine, then you don't need to worry about it. Also, this caveat is only for Stacks and Queues.

## Contribute

Bug reports, PRs, feature requests, all are welcome.

