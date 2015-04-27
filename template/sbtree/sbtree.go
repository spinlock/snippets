package sbtree

type Node struct {
	key   int
	value interface{}

	size        int
	left, right *Node
}

var nilNode = &Node{}

func init() {
	nilNode.left, nilNode.right = nilNode, nilNode
}

func newNode(key int, value interface{}) *Node {
	return &Node{
		key: key, value: value,
		size: 1, left: nilNode, right: nilNode,
	}
}

func (x *Node) Key() int {
	return x.key
}

func (x *Node) Value() interface{} {
	return x.value
}

func (x *Node) check() bool {
	if x.size == 0 {
		return x == nilNode && x == x.left && x == x.right
	} else {
		var pass = true
		if x.left.size != 0 {
			pass = pass && (x.left.key < x.key)
		}
		if x.right.size != 0 {
			pass = pass && (x.key < x.right.key)
		}
		pass = pass && (x.size == x.left.size+x.right.size+1)
		pass = pass && (x.left.size >= x.right.left.size)
		pass = pass && (x.left.size >= x.right.right.size)
		pass = pass && (x.right.size >= x.left.left.size)
		pass = pass && (x.right.size >= x.left.right.size)
		return pass && x.left.check() && x.right.check()
	}
}

type Tree struct {
	root *Node
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) lazyInit() {
	if t.root == nil {
		t.root = nilNode
	}
}

func (t *Tree) Check() bool {
	t.lazyInit()
	return t.root.check()
}

func (t *Tree) Size() int {
	t.lazyInit()
	return t.root.size
}

func (t *Tree) Find(key int) *Node {
	t.lazyInit()
	return t.findByKey(key, t.root)
}

func (t *Tree) Contains(key int) bool {
	t.lazyInit()
	return t.findByKey(key, t.root) != nil
}

func (t *Tree) findByKey(key int, x *Node) *Node {
	for x.size != 0 {
		if x.key == key {
			return x
		} else if x.key < key {
			x = x.right
		} else {
			x = x.left
		}
	}
	return nil
}

func (t *Tree) FindPred(key int) *Node {
	t.lazyInit()
	return t.findPredByKey(key, t.root)
}

func (t *Tree) FindSucc(key int) *Node {
	t.lazyInit()
	return t.findSuccByKey(key, t.root)
}

func (t *Tree) findPredByKey(key int, x *Node) (pred *Node) {
	for x.size != 0 {
		if x.key >= key {
			x = x.left
		} else {
			x, pred = x.right, x
		}
	}
	return pred
}

func (t *Tree) findSuccByKey(key int, x *Node) (succ *Node) {
	for x.size != 0 {
		if x.key <= key {
			x = x.right
		} else {
			x, succ = x.left, x
		}
	}
	return succ
}

func (t *Tree) Rank(key int) int {
	t.lazyInit()
	return t.rank(key, t.root)
}

func (t *Tree) rank(key int, x *Node) (rank int) {
	for x.size != 0 {
		if x.key == key {
			return rank + x.left.size
		} else if x.key < key {
			rank += x.left.size + 1
			x = x.right
		} else {
			x = x.left
		}
	}
	return -(rank + 1)
}

func (t *Tree) Select(rank int) *Node {
	t.lazyInit()
	return t.findByRank(rank, t.root)
}

func (t *Tree) findByRank(rank int, x *Node) *Node {
	if rank < 0 {
		rank += x.size
	}
	for rank >= 0 && rank < x.size {
		if x.left.size == rank {
			return x
		} else if x.left.size < rank {
			rank -= x.left.size + 1
			x = x.right
		} else {
			x = x.left
		}
	}
	return nil
}

func (t *Tree) lrotate(p **Node) {
	x := *p
	y := x.right
	x.right = y.left
	y.left = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *Tree) rrotate(p **Node) {
	x := *p
	y := x.left
	x.left = y.right
	y.right = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *Tree) lbalance(p **Node) {
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

func (t *Tree) rbalance(p **Node) {
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

func (t *Tree) maintain(p **Node) {
	t.lbalance(p)
	t.rbalance(p)
}

func (t *Tree) Insert(key int, value interface{}) (interface{}, bool) {
	t.lazyInit()
	return t.insert(key, value, &t.root)
}

func (t *Tree) insert(key int, value interface{}, p **Node) (oldValue interface{}, addNode bool) {
	x := *p
	if x.size == 0 {
		*p = newNode(key, value)
		return nil, true
	} else if x.key == key {
		oldValue, x.value = x.value, value
		return oldValue, false
	} else if x.key < key {
		oldValue, addNode = t.insert(key, value, &x.right)
	} else {
		oldValue, addNode = t.insert(key, value, &x.left)
	}
	if addNode {
		x.size++
		t.maintain(&x)
		*p = x
	}
	return oldValue, addNode
}

func (t *Tree) Remove(key int) (interface{}, bool) {
	t.lazyInit()
	return t.remove(key, &t.root)
}

func (t *Tree) remove(key int, p **Node) (oldValue interface{}, delNode bool) {
	x := *p
	if x.size == 0 {
		return nil, false
	} else if x.key == key {
		oldValue, delNode = x.value, true
		if x.left.size > x.right.size {
			m := t.findMax(x.left)
			x.key, x.value = m.key, m.value
			t.remove(m.key, &x.left)
		} else if x.right.size != 0 {
			m := t.findMin(x.right)
			x.key, x.value = m.key, m.value
			t.remove(m.key, &x.right)
		}
	} else if x.key < key {
		oldValue, delNode = t.remove(key, &x.right)
	} else {
		oldValue, delNode = t.remove(key, &x.left)
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
	return oldValue, delNode
}

func (t *Tree) findMin(x *Node) *Node {
	for x.left.size != 0 {
		x = x.left
	}
	return x
}

func (t *Tree) findMax(x *Node) *Node {
	for x.right.size != 0 {
		x = x.right
	}
	return x
}

func (t *Tree) Nodes(ns []*Node) []*Node {
	t.lazyInit()
	if n := t.root.size; cap(ns) < n {
		ns = make([]*Node, n)
	} else {
		ns = ns[:n]
	}
	t.nodesWalk(ns, 0, t.root)
	return ns
}

func (t *Tree) nodesWalk(ns []*Node, p int, x *Node) {
	if x.size == 0 {
		return
	}
	t.nodesWalk(ns, p, x.left)
	ns[p+x.left.size] = x
	t.nodesWalk(ns, p+x.left.size+1, x.right)
}
