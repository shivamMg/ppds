package list

import (
  "fmt"
  "reflect"
  "strings"
)

const (
	BoxVer     = "│"
	BoxHor     = "─"
  BoxDownRight = "┌"
  BoxDownLeft = "┐"
	BoxUpRight = "└"
	BoxUpLeft  = "┘"
  BoxVerRight = "├"
  BoxVerLeft = "┤"
)

type Node interface {
  Data() string
  Next() Node
}

func Print(head Node) {
  fmt.Println(Sprint(head))
}

func Sprint(head Node) (s string) {
  // w stores data widths
  data, w := []string{}, []int{}
  for reflect.ValueOf(head) != reflect.Zero(reflect.TypeOf(head)) {
    data = append(data, head.Data())
    w = append(w, len(head.Data()))
    head = head.Next()
  }

  maxDigitWidth := digitWidth(len(data))
  for i := 0; i < len(w); i++ {
    if w[i] < maxDigitWidth {
      w[i] = maxDigitWidth
    }
  }

  for i := 0; i < len(w); i++ {
    s += BoxDownRight + strings.Repeat(BoxHor, w[i]) + BoxDownLeft
  }
  s += "\n"

  s += BoxVer
  for i := 0; i < len(data)-1;i++ {
    s +=  fmt.Sprintf("%*s", w[i], data[i]) + BoxVerRight + BoxVerLeft
  }
  s += fmt.Sprintf("%*s", w[len(w)-1], data[len(data)-1]) + BoxVer + "\n"

  for i := 0; i < len(w); i++ {
    s += BoxVerRight + strings.Repeat(BoxHor, w[i]) + BoxVerLeft
  }
  s += "\n"

  s += BoxVer
  for i := 0; i < len(w)-1; i++ {
    s += fmt.Sprintf("%*d", w[i], i+1) + BoxVer + BoxVer
  }
  s += fmt.Sprintf("%*d", w[len(w)-1], len(w)) + BoxVer + "\n"

  for i := 0; i < len(w); i++ {
    s += BoxUpRight + strings.Repeat(BoxHor, w[i]) + BoxUpLeft
  }
  s += "\n"

  return
}

func digitWidth(d int) (w int) {
	for d != 0 {
		d = d / 10
		w++
	}
	return
}
