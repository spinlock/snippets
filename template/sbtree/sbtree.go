package sbtree

type node struct {
	key   int
	value interface{}

	size        int
	left, right *node
}

var nilNode = &node{}

func init() {
	nilNode.left, nilNode.right = nilNode, nilNode
}

func newNode(key int, value interface{}) *node {
	return &node{
		key: key, value: value,
		size: 1,
		left: nilNode, right: nilNode,
	}
}

func (x *node) check() bool {
	if x.size == 0 {
		if x != nilNode {
			return false
		}
		return x == x.left && x == x.right
	}
	if x.left.size != 0 {
		if x.left.key > x.key {
			return false
		}
	}
	if x.right.size != 0 {
		if x.key > x.right.key {
			return false
		}
	}
	var pass = true
	pass = pass && (x.size == x.left.size+x.right.size+1)
	pass = pass && (x.left.size >= x.right.left.size)
	pass = pass && (x.left.size >= x.right.right.size)
	pass = pass && (x.right.size >= x.left.left.size)
	pass = pass && (x.right.size >= x.left.right.size)
	return pass && x.left.check() && x.right.check()
}

type SBTree struct {
	root *node
}

func New() *SBTree {
	return &SBTree{}
}

func (t *SBTree) lazyInit() {
	if t.root == nil {
		t.root = nilNode
	}
}

func (t *SBTree) Size() int {
	t.lazyInit()
	return t.root.size
}

func (t *SBTree) CheckBalance() bool {
	t.lazyInit()
	return t.root.check()
}

func (t *SBTree) Search(key int) (interface{}, bool) {
	t.lazyInit()
	if x := t.find(key, t.root); x != nil {
		return x.value, true
	}
	return nil, false
}

func (t *SBTree) Contains(key int) bool {
	t.lazyInit()
	return t.find(key, t.root) != nil
}

func (t *SBTree) find(key int, x *node) *node {
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

func (t *SBTree) Rank(key int) int {
	t.lazyInit()
	return t.rank(key, t.root)
}

func (t *SBTree) rank(key int, x *node) (rank int) {
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

func (t *SBTree) Select(rank int) (int, interface{}, bool) {
	t.lazyInit()
	if x := t.findByRank(rank, t.root); x != nil {
		return x.key, x.value, true
	}
	return 0, nil, false
}

func (t *SBTree) findByRank(rank int, x *node) *node {
	if rank < 0 {
		rank += x.size
	}
	for rank < x.size && rank >= 0 {
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

func (t *SBTree) Predecessor(key int) (int, interface{}, bool) {
	t.lazyInit()
	if x := t.findPred(key, t.root); x != nil {
		return x.key, x.value, true
	}
	return 0, nil, false
}

func (t *SBTree) findPred(key int, x *node) (pred *node) {
	for x.size != 0 {
		if x.key >= key {
			x = x.left
		} else {
			x, pred = x.right, x
		}
	}
	return
}

func (t *SBTree) Successor(key int) (int, interface{}, bool) {
	t.lazyInit()
	if x := t.findSucc(key, t.root); x != nil {
		return x.key, x.value, true
	}
	return 0, nil, false
}

func (t *SBTree) findSucc(key int, x *node) (succ *node) {
	for x.size != 0 {
		if x.key <= key {
			x = x.right
		} else {
			x, succ = x.left, x
		}
	}
	return
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

func (t *SBTree) Insert(key int, value interface{}) (interface{}, bool) {
	t.lazyInit()
	return t.insert(key, value, &t.root)
}

func (t *SBTree) insert(key int, value interface{}, p **node) (oldValue interface{}, addNode bool) {
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
	return
}

func (t *SBTree) Delete(key int) (interface{}, bool) {
	t.lazyInit()
	return t.remove(key, &t.root)
}

func (t *SBTree) remove(key int, p **node) (oldValue interface{}, delNode bool) {
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
	return
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

func (t *SBTree) Keys() []int {
	t.lazyInit()
	ks := make([]int, t.root.size)
	t.dumpKeys(ks, 0, t.root)
	return ks
}

func (t *SBTree) dumpKeys(ks []int, n int, x *node) {
	if x.size == 0 {
		return
	}
	t.dumpKeys(ks, n, x.left)
	ks[n+x.left.size] = x.key
	t.dumpKeys(ks, n+x.left.size+1, x.right)
}

func (t *SBTree) Values() []interface{} {
	t.lazyInit()
	vs := make([]interface{}, t.root.size)
	t.dumpValues(vs, 0, t.root)
	return vs
}

func (t *SBTree) dumpValues(vs []interface{}, n int, x *node) {
	if x.size == 0 {
		return
	}
	t.dumpValues(vs, n, x.left)
	vs[n+x.left.size] = x.value
	t.dumpValues(vs, n+x.left.size+1, x.right)
}
