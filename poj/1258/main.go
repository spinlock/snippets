package main

import "fmt"

func process(n int) int {
	var g = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			fmt.Scanf("%d", &g[i][j])
		}
	}

	var d = make([]int, n+1)
	var p = make([]int, n+1)
	var mark = make([]bool, n+1)

	for v := 2; v <= n; v++ {
		if !mark[v] && g[1][v] != 0 {
			d[v] = g[1][v]
			p[v] = 1
		}
	}
	for {
		x := 0
		for v := 2; v <= n; v++ {
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
		for v := 2; v <= n; v++ {
			if !mark[v] && g[x][v] != 0 {
				if p[v] == 0 || d[v] > g[x][v] {
					d[v] = g[x][v]
					p[v] = x
				}
			}
		}
	}

	var sum = 0
	for v := 2; v <= n; v++ {
		sum += d[v]
	}
	return sum
}

func main() {
	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			return
		}
		fmt.Println(process(n))
	}
}
