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

func (x *node) index(key uint8) int {
	beg, end := 0, len(x.keys)-1
	for beg <= end {
		mid := (beg + end) / 2
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

func (x *node) dumpWalk(xs []*node, n int) int {
	if x.tail {
		xs[n] = x
		n = n + 1
	}
	for i := 0; i < len(x.keys); i++ {
		n = x.next[i].dumpWalk(xs, n)
	}
	return n
}

type TrieTree struct {
	root *node
	size int
}

func New() *TrieTree {
	return &TrieTree{}
}

func (t *TrieTree) lazyInit() {
	if t.root == nil {
		t.root = &node{}
	}
}

func (t *TrieTree) Insert(s string) {
	t.lazyInit()
	x := t.root
	for j := 0; j < len(s); j++ {
		b := s[j]
		if i := x.index(b); i >= 0 {
			x = x.next[i]
		} else {
			x = x.insert(b, -(i + 1))
		}
	}
	if !x.tail {
		x.text = s
		x.tail = true
		t.size++
	}
	x.hits++
}

func (t *TrieTree) Dump() []*node {
	xs := make([]*node, t.size)
	t.root.dumpWalk(xs, 0)
	return xs
}

func qswap(xs []*node, i, j int) {
	if i != j {
		xs[i], xs[j] = xs[j], xs[i]
	}
}

func qsort(xs []*node, beg, end int) {
	if beg >= end {
		return
	}
	pivot := beg
	for j := beg + 1; j <= end; j++ {
		if xs[j].text <= xs[beg].text {
			pivot++
			qswap(xs, j, pivot)
		}
	}
	qswap(xs, beg, pivot)
	qsort(xs, beg, pivot-1)
	qsort(xs, pivot+1, end)
}

func main() {
	var r = bufio.NewReader(os.Stdin)
	var t = New()
	for {
		s, _ := r.ReadString('\n')
		if s == "" {
			break
		}
		t.Insert(s[:len(s)-1])
	}
	xs := t.Dump()
	qsort(xs, 0, len(xs)-1)
	total := 0
	for i := 0; i < len(xs); i++ {
		total += xs[i].hits
	}
	for i := 0; i < len(xs); i++ {
		fmt.Printf("%s %0.4f\n", xs[i].text, float64(100*xs[i].hits)/float64(total))
	}
}
