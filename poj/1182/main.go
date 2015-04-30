package main

import "fmt"

type Ufs struct {
	size []int
	dist []int
}

func New(n int) *Ufs {
	ufs := &Ufs{}
	ufs.size = make([]int, n)
	ufs.dist = make([]int, n)
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
		if npx != px {
			ufs.size[x] = npx
			ufs.dist[x] = (ufs.dist[x] + ufs.dist[px]) % 3
		}
		return npx
	}
}

func (ufs *Ufs) Union(x, y int, eat int) bool {
	px := ufs.Find(x)
	py := ufs.Find(y)
	if px == py {
		return ufs.dist[y] == (ufs.dist[x]+eat)%3
	} else {
		ufs.size[px] += ufs.size[py]
		ufs.size[py] = px
		ufs.dist[py] = (ufs.dist[x] - ufs.dist[y] + 3 + eat) % 3
		return true
	}
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	var ufs = New(n + 1)
	var lie = 0
	for i := 0; i < k; i++ {
		var op, x, y int
		fmt.Scanf("%d %d %d", &op, &x, &y)
		var eat int
		if op != 1 {
			eat = 1
		}
		if x > n || y > n {
			lie++
		} else if !ufs.Union(x, y, eat) {
			lie++
		}
	}
	fmt.Println(lie)
}
