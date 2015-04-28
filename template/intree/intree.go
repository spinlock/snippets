package intree

import "math"

const (
	MinInt = int(math.MinInt32)
	MaxInt = int(math.MaxInt32)
)

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
	sum      int
}

type Tree []node

func New(values []int) Tree {
	t := Tree(make([]node, len(values)*4+1))
	t.init(values, 0, 0, len(values)-1)
	return t
}

func (t Tree) init(values []int, i int, beg, end int) {
	t[i].beg = beg
	t[i].end = end
	if beg == end {
		t[i].min = values[beg]
		t[i].max = values[beg]
		t[i].sum = values[beg]
	} else if beg < end {
		mid := beg + (end-beg)/2
		l := i*2 + 1
		r := i*2 + 2
		t.init(values, l, beg, mid)
		t.init(values, r, mid+1, end)
		t[i].min = minInt(t[l].min, t[r].min)
		t[i].max = maxInt(t[l].max, t[r].max)
		t[i].sum = t[l].sum + t[r].sum
	}
}

func (t Tree) Min(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg <= end {
		return t.min(0, beg, end)
	}
	return MaxInt
}

func (t Tree) min(i int, beg, end int) int {
	if t[i].beg == beg && t[i].end == end {
		return t[i].min
	} else {
		mid := t[i].beg + (t[i].end-t[i].beg)/2
		l := i*2 + 1
		r := i*2 + 2
		if end <= mid {
			return t.min(l, beg, end)
		} else if beg > mid {
			return t.min(r, beg, end)
		} else {
			minl := t.min(l, beg, mid)
			minr := t.min(r, mid+1, end)
			return minInt(minl, minr)
		}
	}
}

func (t Tree) Max(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg <= end {
		return t.max(0, beg, end)
	}
	return MinInt
}

func (t Tree) max(i int, beg, end int) int {
	if t[i].beg == beg && t[i].end == end {
		return t[i].max
	} else {
		mid := t[i].beg + (t[i].end-t[i].beg)/2
		l := i*2 + 1
		r := i*2 + 2
		if end <= mid {
			return t.max(l, beg, end)
		} else if beg > mid {
			return t.max(r, beg, end)
		} else {
			maxl := t.max(l, beg, mid)
			maxr := t.max(r, mid+1, end)
			return maxInt(maxl, maxr)
		}
	}
}

func (t Tree) Sum(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg <= end {
		return t.sum(0, beg, end)
	}
	return 0
}

func (t Tree) sum(i int, beg, end int) int {
	if t[i].beg == beg && t[i].end == end {
		return t[i].sum
	} else {
		mid := t[i].beg + (t[i].end-t[i].beg)/2
		l := i*2 + 1
		r := i*2 + 2
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
