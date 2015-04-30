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
	sum, add int
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
		t[x].sum = values[beg]
	} else {
		mid := beg + (end-beg)/2
		l := x*2 + 1
		r := x*2 + 2
		t.init(values, l, beg, mid)
		t.init(values, r, mid+1, end)
		t[x].sum = t[l].sum + t[r].sum
	}
}

func (t Tree) Sum(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end {
		return 0
	}
	return t.sum(0, beg, end)
}

func (t Tree) sum(x int, beg, end int) int {
	if t[x].beg == beg && t[x].end == end {
		return t[x].sum
	} else {
		mid := t[x].beg + (t[x].end-t[x].beg)/2
		l := x*2 + 1
		r := x*2 + 2
		if t[x].add != 0 {
			t[l].add += t[x].add
			t[l].sum += t[x].add * (t[l].end - t[l].beg + 1)
			t[r].add += t[x].add
			t[r].sum += t[x].add * (t[r].end - t[r].beg + 1)
			t[x].add = 0
		}
		if end <= mid {
			return t.sum(l, beg, end)
		} else if beg > mid {
			return t.sum(r, beg, end)
		} else {
			suml := t.sum(l, beg, mid)
			sumr := t.sum(r, mid+1, end)
			return suml + sumr
		}
	}
}

func (t Tree) Add(beg, end int, add int) {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end {
		return
	}
	t.add(0, beg, end, add)
}

func (t Tree) add(x int, beg, end int, add int) {
	t[x].sum += (end - beg + 1) * add
	if t[x].beg == beg && t[x].end == end {
		t[x].add += add
	} else {
		mid := t[x].beg + (t[x].end-t[x].beg)/2
		l := x*2 + 1
		r := x*2 + 2
		if end <= mid {
			t.add(l, beg, end, add)
		} else if beg > mid {
			t.add(r, beg, end, add)
		} else {
			t.add(l, beg, mid, add)
			t.add(r, mid+1, end, add)
		}
	}
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	var values = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &values[i])
	}

	var t = New(values)
	for i := 0; i < m; i++ {
		var op byte
		var beg, end, add int
		fmt.Scanf("%c", &op)
		if op == 'Q' {
			fmt.Scanf("%d %d", &beg, &end)
			fmt.Println(t.Sum(beg-1, end-1))
		} else if op == 'C' {
			fmt.Scanf("%d %d %d", &beg, &end, &add)
			t.Add(beg-1, end-1, add)
		}
	}
}
