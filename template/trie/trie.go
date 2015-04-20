package trie

type node struct {
	keys []uint8
	next []*node
	tail bool
}

func (x *node) find(key uint8) int {
	beg, end := 0, len(x.keys)-1
	for beg <= end {
		mid := (beg + end) / 2
		if x.keys[mid] == key {
			return mid
		} else if x.keys[mid] < key {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -(beg + 1)
}

func (x *node) check() bool {
	for i := 1; i < len(x.keys); i++ {
		if x.keys[i] <= x.keys[i-1] {
			return false
		}
	}
	return true
}

func (x *node) insert(key uint8, i int) *node {
	x.keys = append(x.keys, 0)
	x.next = append(x.next, nil)
	for j := len(x.keys) - 1; j != i; j-- {
		x.keys[j] = x.keys[j-1]
		x.next[j] = x.next[j-1]
	}
	x.keys[i] = key
	x.next[i] = &node{}
	return x.next[i]
}

type TrieTree struct {
	root *node
}

func New() *TrieTree {
	return &TrieTree{}
}

func (t *TrieTree) lazyInit() {
	if t.root == nil {
		t.root = &node{}
	}
}

func (t *TrieTree) Check() bool {
	t.lazyInit()
	return t.root.check()
}

func (t *TrieTree) Insert(s string) {
	t.lazyInit()
	x := t.root
	for j := 0; j < len(s); j++ {
		b := s[j]
		if i := x.find(b); i >= 0 {
			x = x.next[i]
		} else {
			x = x.insert(b, -(i + 1))
		}
	}
	x.tail = true
}

func (t *TrieTree) Contains(s string) bool {
	t.lazyInit()
	x := t.root
	for j := 0; j < len(s); j++ {
		b := s[j]
		if i := x.find(b); i >= 0 {
			x = x.next[i]
		} else {
			return false
		}
	}
	return x.tail
}
