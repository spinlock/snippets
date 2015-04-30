package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	keys []uint8
	next []*node
	tail bool

	text string
	hits int
}

func (x *node) find(key uint8) int {
	beg, end := 0, len(x.keys)-1
	for beg <= end {
		mid := beg + (end-beg)/2
		if x.keys[mid] == key {
			return mid
		} else if x.keys[mid] < key {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -(beg + 1)
}

func (x *node) insert(key uint8, i int) *node {
	x.keys = append(x.keys, 0)
	x.next = append(x.next, nil)
	for j := len(x.keys) - 1; j != i; j-- {
		x.keys[j] = x.keys[j-1]
		x.next[j] = x.next[j-1]
	}
	x.keys[i] = key
	x.next[i] = &node{}
	return x.next[i]
}

type Tree struct {
	root *node
	hits int
	size int
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) lazyInit() {
	if t.root == nil {
		t.root = &node{}
	}
}

func (t *Tree) Insert(s string) {
	t.lazyInit()
	x := t.root
	for i := 0; i < len(s); i++ {
		b := s[i]
		if idx := x.find(b); idx >= 0 {
			x = x.next[idx]
		} else {
			x = x.insert(b, -(idx + 1))
		}
	}
	if !x.tail {
		x.tail = true
		x.text = s
		t.size++
	}
	x.hits++
	t.hits++
}

func (t *Tree) Nodes() []*node {
	t.lazyInit()
	ns := make([]*node, t.size)
	t.nodesWalk(ns, 0, t.root)
	return ns
}

func (t *Tree) nodesWalk(ns []*node, p int, x *node) int {
	if x.tail {
		ns[p] = x
		p++
	}
	for i := 0; i < len(x.next); i++ {
		p = t.nodesWalk(ns, p, x.next[i])
	}
	return p
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
	for i < j && isSpace(s[j-1]) {
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
	var t = New()
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		s = trimSpace(s)
		if len(s) == 0 {
			continue
		}
		t.Insert(s)
	}
	for _, x := range t.Nodes() {
		fmt.Printf("%s %.4f\n", x.text, float64(x.hits)*100/float64(t.hits))
	}
}
