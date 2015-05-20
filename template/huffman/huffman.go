package huffman

type trieTree struct {
	root *trieNode
	size int
}

type trieNode struct {
	nodes []*trieNode
	text  string
	hits  int
	leaf  bool
}

func (n *trieNode) Leaves(ns []*trieNode, i int) int {
	var cnt int
	if n.leaf {
		ns[i] = n
		cnt++
	}
	for _, x := range n.nodes {
		if x != nil {
			cnt += x.Leaves(ns, i+cnt)
		}
	}
	return cnt
}

func (t *trieTree) Insert(s string) {
	if t.root == nil {
		t.root = &trieNode{}
		t.size = 0
	}
	var x = t.root
	for i := 0; i < len(s); i++ {
		b := s[i]
		if x.nodes == nil {
			x.nodes = make([]*trieNode, 256)
		}
		if x.nodes[b] == nil {
			x.nodes[b] = &trieNode{}
		}
		x = x.nodes[b]
	}
	if !x.leaf {
		x.leaf = true
		x.text = s
		t.size++
	}
	x.hits++
}

func (t *trieTree) Nodes() []*huffmanNode {
	if t.size == 0 {
		return nil
	}
	ns := make([]*trieNode, t.size)
	t.root.Leaves(ns, 0)
	hs := make([]*huffmanNode, len(ns))
	for i := 0; i < len(ns); i++ {
		hs[i] = &huffmanNode{text: ns[i].text, weight: ns[i].hits, leaf: true}
	}
	return hs
}

type huffmanNode struct {
	text   string
	leaf   bool
	weight int

	left, right *huffmanNode
}

type huffmanHeap []*huffmanNode

func (h huffmanHeap) reinit() {
	for i := len(h) / 2; i >= 0; i-- {
		h.down(i)
	}
}

func (h huffmanHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h huffmanHeap) down(p int) {
	for p < len(h) {
		l := p*2 + 1
		r := p*2 + 2
		m := p
		if l < len(h) && h[l].weight < h[m].weight {
			m = l
		}
		if r < len(h) && h[r].weight < h[m].weight {
			m = r
		}
		if m == p {
			return
		}
		h.swap(m, p)
		p = m
	}
}

func (h *huffmanHeap) Merge() bool {
	if len(*h) < 2 {
		return false
	}
	var n = len(*h) - 1
	var t = (*h)[0]
	h.swap(0, n)
	*h = (*h)[:n]
	h.down(0)
	var x = &huffmanNode{left: t, right: (*h)[0]}
	x.weight = x.left.weight + x.right.weight
	(*h)[0] = x
	h.down(0)
	return true
}

func (n *huffmanNode) Code(m map[string]string, prefix string) {
	if n.leaf {
		m[n.text] = prefix
	} else {
		n.left.Code(m, prefix+"0")
		n.right.Code(m, prefix+"1")
	}
}

func Encode(ss []string) map[string]string {
	t := &trieTree{}
	for _, s := range ss {
		t.Insert(s)
	}
	var h huffmanHeap = t.Nodes()
	h.reinit()
	for {
		if h.Merge() {
			continue
		}
		break
	}
	var m = make(map[string]string)
	h[0].Code(m, "")
	return m
}
