package main

import "fmt"

type ufs []int

func (t *ufs) init(n int) {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = -1
	}
	*t = array
}

func (t ufs) find(x int) int {
	if px := t[x]; px < 0 {
		return x
	} else {
		t[x] = t.find(px)
		return t[x]
	}
}

func (t ufs) size(x int) int {
	px := t.find(x)
	return -t[px]
}

func (t ufs) union(x, y int) {
	px := t.find(x)
	py := t.find(y)
	if px != py {
		t[px] += t[py]
		t[py] = px
	}
}

func process(n, m int) {
	var t ufs
	t.init(n)
	for i := 0; i < m; i++ {
		var k, x, y int
		fmt.Scanf("%d %d", &k, &x)
		for j := 1; j < k; j++ {
			fmt.Scanf("%d", &y)
			t.union(x, y)
		}
	}
	fmt.Println(t.size(0))
}

func main() {
	var n, m int
	for {
		fmt.Scanf("%d %d", &n, &m)
		if n == 0 && m == 0 {
			return
		}
		process(n, m)
	}
}
