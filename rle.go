package main

import (
	"fmt"
)

func main() {
	s := "0001010111000100000011111010010100011111000"
	z := encode(s)
	fmt.Println(z)
}

func encode(in string) string {
	done := make(chan int)
	c := make(chan string)
	m := Machine{0, 0, c, &done}

	go func() {
		for _, b := range []byte(in) {
			m.Feed(b)
		}
		m.Feed('*')
	}()

	out := ""
	for s := range c {
		if s == "*" {
			break
		}
		out = out + s
	}

	return out
}

type Machine struct {
	B      byte
	N      int
	Output chan string
	Done   *chan int
}

func (m *Machine) Feed(b byte) {
	if b != m.B {
		defer func() {
			m.N = 1
			m.B = b
		}()
		switch m.N {
		case 0:
		case 1:
			m.Output <- string(m.B)
		default:
			m.Output <- fmt.Sprintf("%d%s", m.N, string(m.B))
		}
		if b == '*' {
			m.Output <- "*"
		}
	} else {
		m.N = m.N + 1
	}
}
