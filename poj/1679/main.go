package main

import "fmt"

func process(n, m int) {
	var g = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		var u, v, d int
		fmt.Scanf("%d %d %d", &u, &v, &d)
		g[u][v] = d
		g[v][u] = d
	}

	var mark = make([]bool, n+1)
	var d = make([]int, n+1)
	var p = make([]int, n+1)
	mark[1] = true
	for v := 1; v <= n; v++ {
		if !mark[v] && g[1][v] != 0 {
			d[v] = g[1][v]
			p[v] = 1
		}
	}
	for {
		var x = 0
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
		for v := 1; v <= n; v++ {
			if mark[v] && g[v][x] != 0 {
				if v != p[x] && g[v][x] <= d[x] {
					fmt.Println("Not Unique!")
					return
				}
			}
		}
		mark[x] = true
		for v := 1; v <= n; v++ {
			if !mark[v] && g[x][v] != 0 {
				if p[v] == 0 || d[v] > g[x][v] {
					d[v] = g[x][v]
					p[v] = x
				}
			}
		}
	}

	var sum = 0
	for v := 1; v <= n; v++ {
		sum += d[v]
	}
	fmt.Println(sum)
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var m, n int
		fmt.Scanf("%d %d", &n, &m)
		process(n, m)
	}
}
