package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	code byte
	hits int
	leaf bool

	left, right *node
}

type heap struct {
	buff []*node
	size int
}

func (h *heap) less(i, j int) bool {
	return h.buff[i].hits < h.buff[j].hits
}

func (h *heap) swap(i, j int) {
	if i != j {
		h.buff[i], h.buff[j] = h.buff[j], h.buff[i]
	}
}

func (h *heap) rebuild() {
	for i := h.size / 2; i >= 0; i-- {
		h.down(i)
	}
}

func (h *heap) down(p int) {
	for p < h.size {
		l := p*2 + 1
		r := p*2 + 2
		m := p
		if l < h.size && h.less(l, m) {
			m = l
		}
		if r < h.size && h.less(r, m) {
			m = r
		}
		if p == m {
			return
		}
		h.swap(p, m)
		p = m
	}
}

func (h *heap) up(i int) {
	for i != 0 {
		p := (i - 1) / 2
		if h.less(p, i) {
			return
		}
		h.swap(i, p)
		i = p
	}
}

func (h *heap) pop() *node {
	if h.size == 0 {
		return nil
	}
	n := h.buff[0]
	h.size--
	if h.size != 0 {
		h.swap(0, h.size)
		h.down(0)
	}
	return n
}

func (h *heap) push(n *node) {
	if h.size < len(h.buff) {
		h.buff[h.size] = n
	} else {
		h.buff = append(h.buff, n)
	}
	h.size++
	h.up(h.size - 1)
}

func visit(depth int, n *node) int {
	if n == nil {
		return 0
	}
	if n.leaf {
		return n.hits * depth
	}
	var sum = 0
	sum += visit(depth+1, n.left)
	sum += visit(depth+1, n.right)
	return sum
}

func process(b []byte) {
	var nodes [256]node
	for i := 0; i < len(b); i++ {
		nodes[int(b[i])].hits++
	}
	var h = &heap{}
	for i := 0; i < len(nodes); i++ {
		n := &nodes[i]
		if n.hits != 0 {
			n.code = byte(i)
			n.leaf = true
			h.buff = append(h.buff, n)
			h.size++
		}
	}
	h.rebuild()
	for h.size != 1 {
		a := h.pop()
		b := h.pop()
		n := &node{
			hits: a.hits + b.hits,
			left: a, right: b,
		}
		h.push(n)
	}
	root := h.pop()
	base := len(b) * 8
	var total int
	if root.leaf {
		total = len(b)
	} else {
		total = visit(0, root)
	}
	fmt.Printf("%d %d %.1f\n", base, total, float64(base)/float64(total))
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			return
		}
		b = b[:len(b)-1]
		if string(b) == "END" {
			return
		}
		process(b)
	}
}
