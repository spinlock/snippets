package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	next []*node
	tail bool
}

func (x *node) get(key uint8) *node {
	i := int(key - '0')
	if x.next == nil {
		x.next = make([]*node, 2)
	}
	if x.next[i] == nil {
		x.next[i] = &node{}
	}
	return x.next[i]
}

type Tree struct {
	root *node
}

func (t *Tree) lazyInit() {
	if t.root == nil {
		t.root = &node{}
	}
}

func (t *Tree) InsertNoPrefix(s string) bool {
	t.lazyInit()
	x := t.root
	for i := 0; i < len(s); i++ {
		b := s[i]
		if x.tail {
			return false
		} else {
			x = x.get(b)
		}
	}
	if x.tail || len(x.next) != 0 {
		return false
	} else {
		x.tail = true
		return true
	}
}

func isSpace(b byte) bool {
	const sep = " \t\r\n"
	for i := 0; i < len(sep); i++ {
		if b == sep[i] {
			return true
		}
	}
	return false
}

func trimSpace(s string) string {
	var i, j = 0, len(s)
	for i < j && isSpace(s[i]) {
		i++
	}
	for j > i && isSpace(s[j-1]) {
		j--
	}
	if i < j {
		return s[i:j]
	} else {
		return ""
	}
}

func main() {
	var r = bufio.NewReader(os.Stdin)
	for i := 1; ; i++ {
		var decodable = true
		var t = &Tree{}
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				return
			}
			s = trimSpace(s)
			if len(s) == 0 {
				continue
			}
			if s[0] == '9' {
				break
			}
			if !decodable {
				continue
			}
			decodable = t.InsertNoPrefix(s)
		}
		if decodable {
			fmt.Printf("Set %d is immediately decodable\n", i)
		} else {
			fmt.Printf("Set %d is not immediately decodable\n", i)
		}
	}
}
