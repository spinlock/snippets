package main

import "fmt"

type edge struct {
	u, v int
	cost int
}

func eswap(edges []*edge, i, j int) {
	if i != j {
		edges[i], edges[j] = edges[j], edges[i]
	}
}

func esort(edges []*edge, beg, end int) {
	if beg >= end {
		return
	}
	pivot := beg
	for i := beg + 1; i <= end; i++ {
		if edges[i].cost > edges[beg].cost {
			pivot++
			eswap(edges, i, pivot)
		}
	}
	eswap(edges, beg, pivot)
	esort(edges, beg, pivot-1)
	esort(edges, pivot+1, end)
}

func ufs_init(n int) []int {
	ufs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ufs[i] = -1
	}
	return ufs
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
	if px == py {
		return false
	} else {
		ufs[px] += ufs[py]
		ufs[py] = px
		return true
	}
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	var edges = make([]*edge, m)
	for i := 0; i < m; i++ {
		e := &edge{}
		edges[i] = e
		fmt.Scanf("%d %d %d", &e.u, &e.v, &e.cost)
	}

	esort(edges, 0, m-1)

	var ufs = ufs_init(n)
	var costs int
	for i := 0; i < m; i++ {
		e := edges[i]
		if ufs_join(ufs, e.u, e.v) {
			costs += e.cost
		}
	}
	for i := 2; i <= n; i++ {
		if ufs_find(ufs, i) != ufs_find(ufs, 1) {
			costs = -1
		}
	}
	fmt.Println(costs)
}
