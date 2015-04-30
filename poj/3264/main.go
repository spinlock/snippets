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

type Tree []node

func New(values []int) Tree {
	var t Tree = make([]node, len(values)*4+1)
	if len(values) != 0 {
		t.init(values, 0, 0, len(values)-1)
	}
	return t
}

func (t Tree) init(values []int, x int, beg, end int) {
	t[x].beg = beg
	t[x].end = end
	if beg == end {
		t[x].min = values[beg]
		t[x].max = values[beg]
	} else {
		mid := beg + (end-beg)/2
		l := x*2 + 1
		r := x*2 + 2
		t.init(values, l, beg, mid)
		t.init(values, r, mid+1, end)
		t[x].min = minInt(t[l].min, t[r].min)
		t[x].max = maxInt(t[l].max, t[r].max)
	}
}

func (t Tree) Delta(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end {
		return 0
	}
	min, max := t.minmax(0, beg, end)
	return max - min
}

func (t Tree) minmax(x int, beg, end int) (int, int) {
	if t[x].beg == beg && t[x].end == end {
		return t[x].min, t[x].max
	} else {
		mid := t[x].beg + (t[x].end-t[x].beg)/2
		l := x*2 + 1
		r := x*2 + 2
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
