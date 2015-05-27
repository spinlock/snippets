package main

import "fmt"

func toposort(n int, g [][]int, u int, order, visit []int) {
	if visit[u] != 0 {
		return
	}
	visit[u] = -1
	var max int
	for v := 0; v < n; v++ {
		if g[u][v] == 0 {
			continue
		}
		toposort(n, g, v, order, visit)
		if visit[v] < 0 {
			return
		}
		d := order[v] + 1
		if max < d {
			max = d
		}
	}
	visit[u] = 1
	order[u] = max
}

func process(n int, g [][]int, order, visit []int) int {
	for u := 0; u < n; u++ {
		order[u] = 0
		visit[u] = 0
	}
	for u := 0; u < n; u++ {
		toposort(n, g, u, order, visit)
		if visit[u] < 0 {
			return -1
		}
	}
	for u := 0; u < n; u++ {
		for v := u + 1; v < n; v++ {
			if order[u] == order[v] {
				return 0
			}
		}
	}
	return 1
}

func main() {
	for {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		if n == 0 || m == 0 {
			return
		}
		var g = make([][]int, n)
		for i := 0; i < n; i++ {
			g[i] = make([]int, n)
		}
		var visit = make([]int, n)
		var order = make([]int, n)
		var state int
		for i := 0; i < m; i++ {
			var a, b byte
			fmt.Scanf("%c<%c\n", &a, &b)
			if state != 0 {
				continue
			}
			var x, y = int(a - 'A'), int(b - 'A')
			if g[y][x] != 0 {
				continue
			} else {
				g[y][x] = 1
			}
			s := process(n, g, order, visit)
			if s < 0 {
				fmt.Printf("Inconsistency found after %d relations.\n", i+1)
				state = -1
			} else if s != 0 {
				state = i + 1
			}
		}
		if state == 0 {
			fmt.Println("Sorted sequence cannot be determined.")
		} else if state > 0 {
			fmt.Printf("Sorted sequence determined after %d relations: ", state)
			for i := 0; i < n; i++ {
				for u := 0; u < n; u++ {
					if order[u] == i {
						fmt.Printf("%c", byte(u)+'A')
					}
				}
			}
			fmt.Println(".")
		}
	}
}
