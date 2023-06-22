package tree

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// PrintHr prints the horizontal formatted tree to standard output.
func PrintHr(root Node) {
	fmt.Print(SprintHr(root))
}

// SprintHr returns the horizontal formatted tree.
func SprintHr(root Node) (s string) {
	for _, line := range lines(root) {
		// ignore runes before root node
		line = string([]rune(line)[2:])
		s += strings.TrimRight(line, " ") + "\n"
	}
	return
}

func lines(root Node) (s []string) {
	data := fmt.Sprintf("%s %v ", BoxHor, root.Data())
	l := len(root.Children())
	if l == 0 {
		s = append(s, data)
		return
	}

	w := runeCountInStringExcludingShellEscapeSequences(data)
	for i, c := range root.Children() {
		for j, line := range lines(c) {
			if i == 0 && j == 0 {
				if l == 1 {
					s = append(s, data+BoxHor+line)
				} else {
					s = append(s, data+BoxDownHor+line)
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

// runeCountInStringExcludingShellEscapeSequences is like utf8.RuneCountInString
// but it ignores shell escape sequences.
func runeCountInStringExcludingShellEscapeSequences(s string) (n int) {
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r == '\033' {
			// skip shell escape sequences
			for {
				r, size = utf8.DecodeRuneInString(s)
				if r == 'm' {
					break
				}
				s = s[size:]
			}
		} else {
			n++
		}
		s = s[size:]
	}
	return
}

// PrintHrn prints the horizontal-newline formatted tree to standard output.
func PrintHrn(root Node) {
	fmt.Print(SprintHrn(root))
}

// SprintHrn returns the horizontal-newline formatted tree.
func SprintHrn(root Node) (s string) {
	return strings.Join(lines2(root), "\n") + "\n"
}

func lines2(root Node) (s []string) {
	s = append(s, fmt.Sprintf("%v", root.Data()))
	l := len(root.Children())
	if l == 0 {
		return
	}

	for i, c := range root.Children() {
		for j, line := range lines2(c) {
			// first line of the last child
			if i == l-1 && j == 0 {
				s = append(s, BoxUpRight+BoxHor+" "+line)
			} else if j == 0 {
				s = append(s, BoxVerRight+BoxHor+" "+line)
			} else if i == l-1 {
				s = append(s, "   "+line)
			} else {
				s = append(s, BoxVer+"  "+line)
			}
		}
	}
	return
}
