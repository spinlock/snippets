package intree

import "math"

var MinInt = int(math.MinInt32)
var MaxInt = int(math.MaxInt32)

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

	min, max, sum int
}

type InTree []node

func New(values []int) InTree {
	var t InTree = make([]node, len(values)*4+1)
	if len(values) != 0 {
		t.init(0, 0, len(values)-1, values)
	} else {
		t[0].end = t[0].beg - 1
	}
	return t
}

func (t InTree) init(x int, beg, end int, values []int) {
	t[x].beg = beg
	t[x].end = end
	if beg == end {
		t[x].min = values[beg]
		t[x].max = values[beg]
		t[x].sum = values[beg]
	} else {
		mid := (beg + end) / 2
		l := 2*x + 1
		r := 2*x + 2
		t.init(l, beg, mid, values)
		t.init(r, mid+1, end, values)
		t[x].min = minInt(t[l].min, t[r].min)
		t[x].max = maxInt(t[l].max, t[r].max)
		t[x].sum = t[l].sum + t[r].sum
	}
}

func (t InTree) Min(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end {
		return MaxInt
	}
	return t.min(0, beg, end)
}

func (t InTree) min(x int, beg, end int) int {
	if beg == t[x].beg && end == t[x].end {
		return t[x].min
	}
	mid := (t[x].beg + t[x].end) / 2
	if end <= mid {
		return t.min(2*x+1, beg, end)
	} else if beg >= mid+1 {
		return t.min(2*x+2, beg, end)
	} else {
		minl := t.min(2*x+1, beg, mid)
		minr := t.min(2*x+2, mid+1, end)
		return minInt(minl, minr)
	}
}

func (t InTree) Max(beg, end int) int {
	beg = maxInt(beg, t[0].beg)
	end = minInt(end, t[0].end)
	if beg > end {
		return MinInt
	}
	return t.max(0, beg, end)
}

func (t InTree) max(x int, beg, end int) int {
	if beg == t[x].beg && end == t[x].end {
		return t[x].max
	}
	mid := (t[x].beg + t[x].end) / 2
	if end <= mid {
		return t.max(2*x+1, beg, end)
	} else if beg >= mid+1 {
		return t.max(2*x+2, beg, end)
	} else {
		maxl := t.max(2*x+1, beg, mid)
		maxr := t.max(2*x+2, mid+1, end)
		return maxInt(maxl, maxr)
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
	}
	mid := (t[x].beg + t[x].end) / 2
	if end <= mid {
		return t.sum(2*x+1, beg, end)
	} else if beg >= mid+1 {
		return t.sum(2*x+2, beg, end)
	} else {
		suml := t.sum(2*x+1, beg, mid)
		sumr := t.sum(2*x+2, mid+1, end)
		return suml + sumr
	}
}
