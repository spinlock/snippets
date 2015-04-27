package ufs

type Ufs []int

func New(n int) Ufs {
	ufs := make([]int, n)
	for i := 0; i < n; i++ {
		ufs[i] = -1
	}
	return ufs
}

func (ufs Ufs) Find(x int) int {
	if px := ufs[x]; px < 0 {
		return x
	} else {
		ufs[x] = ufs.Find(px)
		return ufs[x]
	}
}

func (ufs Ufs) Size(x int) int {
	px := ufs.Find(x)
	return -ufs[px]
}

func (ufs Ufs) Union(x, y int) {
	px := ufs.Find(x)
	py := ufs.Find(y)
	if px != py {
		ufs[px] += ufs[py]
		ufs[py] = px
	}
}
