package main

import "fmt"

func minInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}

func maxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

type node struct {
	beg, end int
	min, max int
}

type Tree struct {
	nodes []node
}

func New(values []int) *Tree {
	t := &Tree{}
	t.nodes = make([]node, len(values)*4+1)
	if len(values) != 0 {
		t.init(values, 0, 0, len(values)-1)
	}
	return t
}

func (t *Tree) init(values []int, p int, beg, end int) {
	x := &t.nodes[p]
	x.beg = beg
	x.end = end
	if beg == end {
		x.min = values[beg]
		x.max = values[beg]
	} else {
		mid := beg + (end-beg)/2
		l := p*2 + 1
		r := p*2 + 2
		t.init(values, l, beg, mid)
		t.init(values, r, mid+1, end)
		x.min = minInt(t.nodes[l].min, t.nodes[r].min)
		x.max = maxInt(t.nodes[l].max, t.nodes[r].max)
	}
}

func (t *Tree) Delta(beg, end int) int {
	beg = maxInt(beg, t.nodes[0].beg)
	end = minInt(end, t.nodes[0].end)
	if beg > end {
		return 0
	}
	min, max := t.minmax(0, beg, end)
	return max - min
}

func (t *Tree) minmax(p int, beg, end int) (int, int) {
	x := &t.nodes[p]
	if x.beg == beg && x.end == end {
		return x.min, x.max
	} else {
		mid := x.beg + (x.end-x.beg)/2
		l := p*2 + 1
		r := p*2 + 2
		if end <= mid {
			return t.minmax(l, beg, end)
		} else if beg > mid {
			return t.minmax(r, beg, end)
		} else {
			minl, maxl := t.minmax(l, beg, mid)
			minr, maxr := t.minmax(r, mid+1, end)
			return minInt(minl, minr), maxInt(maxl, maxr)
		}
	}
}

func main() {
	var m, n int
	fmt.Scanf("%d %d", &n, &m)

	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &values[i])
	}

	var t = New(values)
	for i := 0; i < m; i++ {
		var beg, end int
		fmt.Scanf("%d %d", &beg, &end)
		fmt.Println(t.Delta(beg-1, end-1))
	}
}
