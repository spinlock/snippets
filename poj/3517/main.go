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

type Tree struct {
	root *node
}

func New(n int) *Tree {
	var t = &Tree{}
	t.root = t.build(0, n-1)
	return t
}

func (t *Tree) build(beg, end int) *node {
	if beg > end {
		return nilNode
	} else {
		mid := beg + (end-beg)/2
		return &node{
			key:   mid,
			size:  end - beg + 1,
			left:  t.build(beg, mid-1),
			right: t.build(mid+1, end),
		}
	}
}

func (t *Tree) Size() int {
	return t.root.size
}

func (t *Tree) DeleteRank(rank int) (int, bool) {
	return t.removeRank(rank, &t.root)
}

func (t *Tree) removeRank(rank int, p **node) (delKey int, delNode bool) {
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

func (t *Tree) findMin(x *node) *node {
	for x.left.size != 0 {
		x = x.left
	}
	return x
}

func (t *Tree) findMax(x *node) *node {
	for x.right.size != 0 {
		x = x.right
	}
	return x
}

func (t *Tree) lrotate(p **node) {
	x := *p
	y := x.right
	x.right = y.left
	y.left = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *Tree) rrotate(p **node) {
	x := *p
	y := x.left
	x.left = y.right
	y.right = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *Tree) lbalance(p **node) {
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

func (t *Tree) rbalance(p **node) {
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

func (t *Tree) maintain(p **node) {
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
		var t = New(n)
		var rank = m - 1
		var last int
		for i := 0; i < n; i++ {
			rank = rank % t.Size()
			last, _ = t.DeleteRank(rank)
			rank = rank + k - 1
		}
		fmt.Println(last + 1)
	}
}
