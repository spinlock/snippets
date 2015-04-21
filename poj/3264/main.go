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

type InTree []node

func NewTree(values []int) InTree {
	var t InTree = make([]node, len(values)*4+1)
	if len(values) != 0 {
		t.init(0, values, 0, len(values)-1)
	} else {
		t[0].end = t[0].beg - 1
	}
	return t
}

func (t InTree) init(x int, values []int, beg, end int) {
	t[x].beg = beg
	t[x].end = end
	if beg == end {
		t[x].min = values[beg]
		t[x].max = values[beg]
	} else {
		mid := (beg + end) / 2
		l := 2*x + 1
		r := 2*x + 2
		t.init(l, values, beg, mid)
		t.init(r, values, mid+1, end)
		t[x].min = minInt(t[l].min, t[r].min)
		t[x].max = maxInt(t[l].max, t[r].max)
	}
}

func (t InTree) MinMax(beg, end int) (int, int) {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end {
		return 0, 0
	}
	return t.minmax(0, beg, end)
}

func (t InTree) minmax(x int, beg, end int) (int, int) {
	if beg == t[x].beg && end == t[x].end {
		return t[x].min, t[x].max
	}
	mid := (t[x].beg + t[x].end) / 2
	l := 2*x + 1
	r := 2*x + 2
	if mid >= end {
		return t.minmax(l, beg, end)
	} else if beg >= mid+1 {
		return t.minmax(r, beg, end)
	} else {
		minl, maxl := t.minmax(l, beg, mid)
		minr, maxr := t.minmax(r, mid+1, end)
		return minInt(minl, minr), maxInt(maxl, maxr)
	}
}

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)
	height := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &height[i])
	}
	var t = NewTree(height)
	for i := 0; i < q; i++ {
		var beg, end int
		fmt.Scanf("%d %d", &beg, &end)
		var min, max = t.MinMax(beg, end)
		fmt.Printf("%d\n", max-min)
	}
}
