package heap

type Heap struct {
	buff []int
	size int
}

func New() *Heap {
	return &Heap{}
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) swap(i, j int) {
	h.buff[i], h.buff[j] = h.buff[j], h.buff[i]
}

func (h *Heap) down(p int) {
	for p < h.size {
		l := p*2 + 1
		r := p*2 + 2
		m := p
		if l < h.size && h.buff[l] < h.buff[m] {
			m = l
		}
		if r < h.size && h.buff[r] < h.buff[m] {
			m = r
		}
		if p == m {
			return
		}
		h.swap(p, m)
		p = m
	}
}

func (h *Heap) up(i int) {
	for i != 0 {
		p := (i - 1) / 2
		if h.buff[p] <= h.buff[i] {
			return
		}
		h.swap(p, i)
		i = p
	}
}

func (h *Heap) Push(v int) {
	if h.size == len(h.buff) {
		h.buff = append(h.buff, v)
	} else {
		h.buff[h.size] = v
	}
	h.size++
	h.up(h.size - 1)
}

func (h *Heap) Pop() (int, bool) {
	if h.size == 0 {
		return 0, false
	}
	v := h.buff[0]
	h.size--
	if h.size != 0 {
		h.swap(0, h.size)
		h.down(0)
	}
	return v, true
}
