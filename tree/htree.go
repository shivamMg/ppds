package tree

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// PrintHr prints the horizontal formatted tree to standard output.
func PrintHr(root Node) {
	fmt.Print(Sprint(root))
}

// SprintHr returns the horizontal formatted tree.
func SprintHr(root Node) string {
	return strings.Join(lines(root), "\n") + "\n"
}

func lines(root Node) (s []string) {
	l := len(root.Children())
	if l == 0 {
		s = append(s, root.Data())
		return
	}

	w := utf8.RuneCountInString(root.Data())
	for i, c := range root.Children() {
		for j, line := range lines(c) {
			if i == 0 && j == 0 {
				if l == 1 {
					s = append(s, root.Data()+BoxHor+line)
				} else {
					s = append(s, root.Data()+BoxDownHor+line)
				}
				continue
			}

			var box string
			if i == l-1 && j == 0 {
				// first line of the last child
				box = BoxUpRight
			} else if i == l-1 {
				box = " "
			} else if j == 0 {
				box = BoxVerRight
			} else {
				box = BoxVer
			}
			s = append(s, strings.Repeat(" ", w)+box+line)
		}
	}
	return
}
