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
		t[x].sum = values[beg]
	} else {
		mid := (beg + end) / 2
		l := 2*x + 1
		r := 2*x + 2
		t.init(l, values, beg, mid)
		t.init(r, values, mid+1, end)
		t[x].sum = t[l].sum + t[r].sum
	}
}

func (t InTree) Sum(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end {
		return 0
	}
	return t.sum(0, beg, end)
}

func (t InTree) sum(x int, beg, end int) int {
	if beg == t[x].beg && end == t[x].end {
		return t[x].sum
	} else {
		mid := (t[x].beg + t[x].end) / 2
		l := 2*x + 1
		r := 2*x + 2
		if t[x].add != 0 {
			t[l].add += t[x].add
			t[r].add += t[x].add
			t[l].sum += (t[l].end - t[l].beg + 1) * t[x].add
			t[r].sum += (t[r].end - t[r].beg + 1) * t[x].add
			t[x].add = 0
		}
		if end <= mid {
			return t.sum(l, beg, end)
		} else if beg >= mid+1 {
			return t.sum(r, beg, end)
		} else {
			suml := t.sum(l, beg, mid)
			sumr := t.sum(r, mid+1, end)
			return suml + sumr
		}
	}
}

func (t InTree) Add(beg, end int, add int) {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end || add == 0 {
		return
	}
	t.add(0, beg, end, add)
}

func (t InTree) add(x int, beg, end int, add int) {
	if beg == t[x].beg && end == t[x].end {
		t[x].add += add
	} else {
		mid := (t[x].beg + t[x].end) / 2
		l := 2*x + 1
		r := 2*x + 2
		if end <= mid {
			t.add(l, beg, end, add)
		} else if beg >= mid+1 {
			t.add(r, beg, end, add)
		} else {
			t.add(l, beg, mid, add)
			t.add(r, mid+1, end, add)
		}
	}
	t[x].sum += (end - beg + 1) * add
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
		var op byte
		fmt.Scanf("%c", &op)
		switch op {
		case 'Q':
			fmt.Scanf("%d %d\n", &beg, &end)
			fmt.Printf("%d\n", t.Sum(beg, end))
		case 'C':
			var add int
			fmt.Scanf("%d %d %d\n", &beg, &end, &add)
			t.Add(beg, end, add)
		}
	}
}
