package main

import "fmt"

const m, n = 44, 1995

func eular(g [][]int, u int, path []int, size int) int {
	for e := 1; e <= n; e++ {
		if v := g[e][u]; v != 0 {
			g[e][u] = 0
			g[e][v] = 0
			size = eular(g, v, path, size)
			path[size] = e
			size = size + 1
		}
	}
	return size
}

func reverse(path []int, i, j int) {
	for i < j {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}
}

func process() int {
	var g = make([][]int, n+1)
	for e := 1; e <= n; e++ {
		g[e] = make([]int, m+1)
	}
	var d = make([]int, m+1)
	for i := 0; ; i++ {
		var x, y, z int
		fmt.Scanf("%d %d", &x, &y)
		if x == 0 || y == 0 {
			if i == 0 {
				return -1
			}
			break
		}
		fmt.Scanf("%d", &z)
		d[x]++
		d[y]++
		g[z][x] = y
		g[z][y] = x
	}

	for u := 1; u <= m; u++ {
		if d[u]%2 != 0 {
			fmt.Println("Round trip does not exist.")
			return 0
		}
	}

	var path = make([]int, n+1)
	size := eular(g, 1, path, 0)
	reverse(path, 0, size-1)
	for i := 0; i < size; i++ {
		fmt.Print(path[i], " ")
	}
	fmt.Println()
	return 0
}

func main() {
	for {
		if process() != 0 {
			return
		}
	}
}
