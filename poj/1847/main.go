package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	var g = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		var m, v int
		fmt.Scanf("%d", &m)
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &v)
			if j == 1 {
				g[i][v] = 0
			} else {
				g[i][v] = 1
			}
		}
	}

	var mark = make([]bool, n+1)
	var d = make([]int, n+1)
	var p = make([]int, n+1)

	mark[a] = true
	p[a] = a
	for v := 1; v <= n; v++ {
		if !mark[v] && g[a][v] >= 0 {
			d[v] = g[a][v]
			p[v] = a
		}
	}

	for {
		var x int
		for v := 1; v <= n; v++ {
			if !mark[v] && p[v] != 0 {
				if x == 0 || d[x] > d[v] {
					x = v
				}
			}
		}
		if x == 0 {
			break
		}
		mark[x] = true
		for v := 1; v <= n; v++ {
			if !mark[v] && g[x][v] >= 0 {
				if p[v] == 0 || d[v] > d[x]+g[x][v] {
					d[v] = d[x] + g[x][v]
					p[v] = x
				}
			}
		}
	}

	if p[b] == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(d[b])
	}
}
