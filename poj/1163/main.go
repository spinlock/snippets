package main

import "fmt"

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var g = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scanf("%d", &g[i][j])
		}
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			g[i][j] += max(g[i+1][j], g[i+1][j+1])
		}
	}
	fmt.Println(g[1][1])
}
