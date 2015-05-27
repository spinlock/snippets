package main

import "fmt"

func dfs(m [][]byte, col []bool, n, k, i int) int {
	if k == 0 {
		return 1
	} else if i > n {
		return 0
	}
	var cnt = 0
	for j := 1; j <= n; j++ {
		if col[j] == false && m[i][j] == '#' {
			col[j] = true
			cnt += dfs(m, col, n, k-1, i+1)
			col[j] = false
		}
	}
	cnt += dfs(m, col, n, k, i+1)
	return cnt
}

func process(n, k int) int {
	var m = make([][]byte, n+1)
	for i := 1; i <= n; i++ {
		m[i] = make([]byte, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Scanf("%c", &m[i][j])
		}
		fmt.Scanf("\n")
	}

	var col = make([]bool, n+1)
	return dfs(m, col, n, k, 1)
}

func main() {
	var n, k int
	for {
		fmt.Scanf("%d %d", &n, &k)
		if n < 0 || k < 0 {
			return
		}
		fmt.Println(process(n, k))
	}
}
