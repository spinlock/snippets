package main

import "fmt"

type node struct {
	key int

	size int

	left, right *node
}

var nilNode = &node{}

func init() {
	nilNode.left, nilNode.right = nilNode, nilNode
}

type SBTree struct {
	root  *node
	nodes []node
}

func NewTree(keys []int) *SBTree {
	var t = &SBTree{}
	t.nodes = make([]node, len(keys))
	t.root = t.rebuild(keys, 0, len(keys)-1)
	return t
}

func (t *SBTree) rebuild(keys []int, beg, end int) *node {
	if beg > end {
		return nilNode
	} else if beg == end {
		t.nodes[beg].key = keys[beg]
		t.nodes[beg].size = 1
		t.nodes[beg].left = nilNode
		t.nodes[beg].right = nilNode
		return &t.nodes[beg]
	} else {
		mid := (beg + end) / 2
		t.nodes[mid].key = keys[mid]
		t.nodes[mid].size = end - beg + 1
		t.nodes[mid].left = t.rebuild(keys, beg, mid-1)
		t.nodes[mid].right = t.rebuild(keys, mid+1, end)
		return &t.nodes[mid]
	}
}

func (t *SBTree) Size() int {
	return t.root.size
}

func (t *SBTree) DeleteRank(rank int) (int, bool) {
	return t.removeRank(rank, &t.root)
}

func (t *SBTree) removeRank(rank int, p **node) (delKey int, delNode bool) {
	x := *p
	if rank < 0 {
		rank += x.size
	}
	if rank < 0 || rank >= x.size {
		return 0, false
	} else if x.left.size == rank {
		delKey, delNode = x.key, true
		if x.left.size > x.right.size {
			m := t.findMax(x.left)
			x.key = m.key
			t.removeRank(-1, &x.left)
		} else if x.right.size != 0 {
			m := t.findMin(x.right)
			x.key = m.key
			t.removeRank(0, &x.right)
		}
	} else if x.left.size < rank {
		rank -= x.left.size + 1
		delKey, delNode = t.removeRank(rank, &x.right)
	} else {
		delKey, delNode = t.removeRank(rank, &x.left)
	}
	if delNode {
		x.size--
		if x.size == 0 {
			*p = nilNode
		} else {
			t.maintain(&x)
			*p = x
		}
	}
	return delKey, delNode
}

func (t *SBTree) findMin(x *node) *node {
	for x.left.size != 0 {
		x = x.left
	}
	return x
}

func (t *SBTree) findMax(x *node) *node {
	for x.right.size != 0 {
		x = x.right
	}
	return x
}

func (t *SBTree) lrotate(p **node) {
	x := *p
	y := x.right
	x.right = y.left
	y.left = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *SBTree) rrotate(p **node) {
	x := *p
	y := x.left
	x.left = y.right
	y.right = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *SBTree) lbalance(p **node) {
	x := *p
	if x.right.size < x.left.left.size {
		t.rrotate(&x)
	} else if x.right.size < x.left.right.size {
		t.lrotate(&x.left)
		t.rrotate(&x)
	} else {
		return
	}
	t.rbalance(&x.right)
	t.lbalance(&x.left)
	t.maintain(&x)
	*p = x
}

func (t *SBTree) rbalance(p **node) {
	x := *p
	if x.left.size < x.right.right.size {
		t.lrotate(&x)
	} else if x.left.size < x.right.left.size {
		t.rrotate(&x.right)
		t.lrotate(&x)
	} else {
		return
	}
	t.lbalance(&x.left)
	t.rbalance(&x.right)
	t.maintain(&x)
	*p = x
}

func (t *SBTree) maintain(p **node) {
	t.lbalance(p)
	t.rbalance(p)
}

func main() {
	for {
		var n, k, m int
		fmt.Scanf("%d %d %d", &n, &k, &m)
		if n == 0 {
			return
		}
		keys := make([]int, n)
		for i := 0; i < n; i++ {
			keys[i] = i + 1
		}
		var t = NewTree(keys)
		var rank = m - 1
		var last int
		for t.Size() != 0 {
			rank = rank % t.Size()
			last, _ = t.DeleteRank(rank)
			rank = rank + k - 1
		}
		fmt.Println(last)
	}
}
