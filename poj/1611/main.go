package main

import "fmt"

type Ufs struct {
	size []int
}

func New(n int) *Ufs {
	ufs := &Ufs{}
	ufs.size = make([]int, n)
	for i := 0; i < n; i++ {
		ufs.size[i] = -1
	}
	return ufs
}

func (ufs *Ufs) Find(x int) int {
	if px := ufs.size[x]; px < 0 {
		return x
	} else {
		npx := ufs.Find(px)
		ufs.size[x] = npx
		return npx
	}
}

func (ufs *Ufs) Size(x int) int {
	px := ufs.Find(x)
	return -ufs.size[px]
}

func (ufs *Ufs) Union(x, y int) {
	px := ufs.Find(x)
	py := ufs.Find(y)
	if px != py {
		ufs.size[px] += ufs.size[py]
		ufs.size[py] = px
	}
}

func main() {
	var n, m int
	for {
		fmt.Scanf("%d %d", &n, &m)
		if n == 0 {
			return
		}
		ufs := New(n)
		for i := 0; i < m; i++ {
			var k, x, y int
			fmt.Scanf("%d %d", &k, &x)
			for j := 1; j < k; j++ {
				fmt.Scanf("%d", &y)
				ufs.Union(x, y)
			}
		}
		fmt.Println(ufs.Size(0))
	}
}
