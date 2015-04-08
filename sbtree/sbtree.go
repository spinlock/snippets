package sbtree

type sbnode struct {
	key   int
	value interface{}

	size        int
	left, right *sbnode
}

func newNode(key int, value interface{}) *sbnode {
	return &sbnode{
		key: key, value: value,
		size: 1,
		left: nilNode, right: nilNode,
	}
}

var nilNode = &sbnode{}

func init() {
	nilNode.left, nilNode.right = nilNode, nilNode
}

type SBTree struct {
	root *sbnode
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

func (t *SBTree) lrotate(p **sbnode) {
	x := *p
	y := x.right
	x.right = y.left
	y.left = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *SBTree) rrotate(p **sbnode) {
	x := *p
	y := x.left
	x.left = y.right
	y.right = x
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
	*p = y
}

func (t *SBTree) lbalance(p **sbnode) {
	x := *p
	if x.right.size < x.left.left.size {
		t.rrotate(&x)
	} else if x.right.size < x.left.right.size {
		t.lrotate(&x.left)
		t.rrotate(&x)
	} else {
		return
	}
	*p = x
	t.rbalance(&x.right)
	t.lbalance(&x.left)
}

func (t *SBTree) rbalance(p **sbnode) {
	x := *p
	if x.left.size < x.right.right.size {
		t.lrotate(&x)
	} else if x.left.size < x.right.left.size {
		t.rrotate(&x.right)
		t.lrotate(&x)
	} else {
		return
	}
	*p = x
	t.lbalance(&x.left)
	t.rbalance(&x.right)
}

func (t *SBTree) maintain(p **sbnode) {
	x := *p
	t.lbalance(&x)
	t.rbalance(&x)
	*p = x
}

func (t *SBTree) Search(key int) (interface{}, bool) {
	t.lazyInit()
	x := t.search(key)
	if x != nil {
		return x.value, true
	}
	return nil, false
}

func (t *SBTree) Contains(key int) bool {
	t.lazyInit()
	return t.search(key) != nil
}

func (t *SBTree) search(key int) *sbnode {
	x := t.root
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
	x, rank := t.root, 0
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
	x := t.root
	if rank < 0 || rank >= x.size {
		return 0, nil, false
	}
	for {
		if x.left.size == rank {
			return x.key, x.value, true
		} else if x.left.size < rank {
			rank -= x.left.size + 1
			x = x.right
		} else {
			x = x.left
		}
	}
}

func (t *SBTree) Predecessor(key int) (int, interface{}, bool) {
	t.lazyInit()
	var p *sbnode
	x := t.root
	for x.size != 0 {
		if x.key >= key {
			x = x.left
		} else {
			x, p = x.right, x
		}
	}
	if p != nil {
		return p.key, p.value, true
	}
	return 0, nil, false
}

func (t *SBTree) Successor(key int) (int, interface{}, bool) {
	t.lazyInit()
	var s *sbnode
	x := t.root
	for x.size != 0 {
		if x.key <= key {
			x = x.right
		} else {
			x, s = x.left, x
		}
	}
	if s != nil {
		return s.key, s.value, true
	}
	return 0, nil, false
}

func (t *SBTree) Insert(key int, value interface{}) (interface{}, bool) {
	t.lazyInit()
	return t.insert(key, value, &t.root)
}

func (t *SBTree) insert(key int, value interface{}, p **sbnode) (oldValue interface{}, addNode bool) {
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
	return t.delete(key, &t.root)
}

func (t *SBTree) delete(key int, p **sbnode) (oldValue interface{}, delNode bool) {
	x := *p
	if x.size == 0 {
		return nil, false
	} else if x.key == key {
		oldValue, delNode = x.value, true
		if x.left.size > x.right.size {
			p := t.findMax(x.left)
			x.key, x.value = p.key, p.value
			t.delete(p.key, &x.left)
		} else if x.right.size != 0 {
			p := t.findMin(x.right)
			x.key, x.value = p.key, p.value
			t.delete(p.key, &x.right)
		}
	} else if x.key < key {
		oldValue, delNode = t.delete(key, &x.right)
	} else {
		oldValue, delNode = t.delete(key, &x.left)
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

func (t *SBTree) findMin(x *sbnode) *sbnode {
	for x.left.size != 0 {
		x = x.left
	}
	return x
}

func (t *SBTree) findMax(x *sbnode) *sbnode {
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

func (t *SBTree) dumpKeys(ks []int, pos int, x *sbnode) {
	if x.size == 0 {
		return
	}
	t.dumpKeys(ks, pos, x.left)
	ks[pos+x.left.size] = x.key
	t.dumpKeys(ks, pos+x.left.size+1, x.right)
}

func (t *SBTree) Values() []interface{} {
	t.lazyInit()
	vs := make([]interface{}, t.root.size)
	t.dumpValues(vs, 0, t.root)
	return vs
}

func (t *SBTree) dumpValues(vs []interface{}, pos int, x *sbnode) {
	if x.size == 0 {
		return
	}
	t.dumpValues(vs, pos, x.left)
	vs[pos+x.left.size] = x.value
	t.dumpValues(vs, pos+x.left.size+1, x.right)
}
