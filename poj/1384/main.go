package main

import "fmt"

func log2(w int) int {
	var p int
	for x := 1; x < w; x *= 2 {
		p = p + 1
	}
	return p
}

func process(m, w int) int {
	var values1 = make([]int, m+1)
	var weight1 = make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Scanf("%d %d", &values1[i], &weight1[i])
	}

	var p = log2(w) + 1
	var values2 = make([]int, m*p+1)
	var weight2 = make([]int, m*p+1)
	var n int
	for i := 1; i <= m; i++ {
		vx, wx := values1[i], weight1[i]
		for wx <= w {
			n++
			values2[n] = vx
			weight2[n] = wx
			vx *= 2
			wx *= 2
		}
	}

	var so = make([]int, w+1)
	var sn = make([]int, w+1)

	for j := 1; j <= w; j++ {
		so[j] = -1
	}
	for i := 1; i <= n; i++ {
		vx, wx := values2[i], weight2[i]
		for j := 1; j <= w; j++ {
			sn[j] = so[j]
			if j >= wx && so[j-wx] >= 0 {
				if sn[j] < 0 || sn[j] > so[j-wx]+vx {
					sn[j] = so[j-wx] + vx
				}
			}
		}
		so, sn = sn, so
	}
	return so[w]
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var e, f, n int
		fmt.Scanf("%d %d", &e, &f)
		fmt.Scanf("%d", &n)
		min := process(n, f-e)
		if min >= 0 {
			fmt.Printf("The minimum amount of money in the piggy-bank is %d.\n", min)
		} else {
			fmt.Printf("This is impossible.\n")
		}
	}
}
