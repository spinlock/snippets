package main

import "fmt"

type edge struct {
	u, v int
	d2   int
}

func eswap(edges []*edge, i, j int) {
	if i != j {
		edges[i], edges[j] = edges[j], edges[i]
	}
}

func ufs_init(ufs []int, n int) {
	for i := 1; i <= n; i++ {
		ufs[i] = -1
	}
}

func ufs_find(ufs []int, x int) int {
	if px := ufs[x]; px < 0 {
		return x
	} else {
		ufs[x] = ufs_find(ufs, px)
		return ufs[x]
	}
}

func ufs_join(ufs []int, x, y int) bool {
	px := ufs_find(ufs, x)
	py := ufs_find(ufs, y)
	if px != py {
		ufs[px] += ufs[py]
		ufs[py] = px
		return true
	} else {
		return false
	}
}

func esort(edges []*edge, beg, end int) {
	if beg >= end {
		return
	}
	pivot := beg
	for i := beg + 1; i <= end; i++ {
		if edges[i].d2 < edges[beg].d2 {
			pivot++
			eswap(edges, i, pivot)
		}
	}
	eswap(edges, beg, pivot)
	esort(edges, beg, pivot-1)
	esort(edges, pivot+1, end)
}

func main() {
	var n, m int
	fmt.Scanf("%d", &n)

	var x = make([]int, n+1)
	var y = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d", &x[i], &y[i])
	}

	var edges = make([]*edge, n*n)
	var k int
	for u := 1; u <= n; u++ {
		for v := u + 1; v <= n; v++ {
			dx := x[u] - x[v]
			dy := y[u] - y[v]
			e := &edge{
				u: u, v: v,
				d2: dx*dx + dy*dy,
			}
			edges[k] = e
			k++
		}
	}
	esort(edges, 0, k-1)

	ufs := make([]int, n+1)
	ufs_init(ufs, n)

	fmt.Scanf("%d", &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Scanf("%d %d", &u, &v)
		ufs_join(ufs, u, v)
	}

	for i := 0; i < k; i++ {
		e := edges[i]
		if ufs_join(ufs, e.u, e.v) {
			fmt.Println(e.u, e.v)
		}
	}
}
