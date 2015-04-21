package main

import "fmt"

type ufs struct {
	array []int
	mod   []int
}

func (t *ufs) init(n int) {
	t.array = make([]int, n)
	for i := 0; i < n; i++ {
		t.array[i] = -1
	}
	t.mod = make([]int, n)
}

func (t *ufs) find(x int) int {
	if px := t.array[x]; px < 0 {
		return x
	} else {
		t.array[x] = t.find(px)
		t.mod[x] = (t.mod[x] + t.mod[px]) % 3
		return t.array[x]
	}
}

func (t *ufs) union(x, y int, eat int) bool {
	px := t.find(x)
	py := t.find(y)
	if px == py {
		return (t.mod[x]+eat)%3 == t.mod[y]
	} else {
		t.array[px] += t.array[py]
		t.array[py] = px
		t.mod[py] = (t.mod[x] + 3 + eat - t.mod[y]) % 3
		return true
	}
}

func check(x int, beg, end int) bool {
	return x >= beg && x <= end
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	var t ufs
	t.init(n + 1)
	var failed int
	for i := 0; i < m; i++ {
		var d, x, y int
		fmt.Scanf("%d %d %d", &d, &x, &y)
		if !check(x, 1, n) || !check(y, 1, n) {
			failed++
		} else if d == 1 {
			if !t.union(x, y, 0) {
				failed++
			}
		} else if d == 2 {
			if !t.union(x, y, 1) {
				failed++
			}
		}
	}
	fmt.Println(failed)
}
