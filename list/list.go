package list

import (
  "fmt"
  "reflect"
)

const (
	BoxVer     = "│"
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

  s += BoxVer
  for i := 0; i < len(data)-1;i++ {
    s +=  data[i] + BoxVerRight + BoxVerLeft
  }
  s += data[len(data)-1] + BoxVer
  return
}
