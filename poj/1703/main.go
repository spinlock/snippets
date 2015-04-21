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
		t.mod[x] = (t.mod[x] + t.mod[px]) % 2
		return t.array[x]
	}
}

func (t *ufs) union(x, y int) bool {
	px := t.find(x)
	py := t.find(y)
	if px == py {
		return t.mod[x] != t.mod[y]
	} else {
		t.array[px] += t.array[py]
		t.array[py] = px
		t.mod[py] = (t.mod[x] + 3 - t.mod[y]) % 2
		return true
	}
}

func (t *ufs) gang(x, y int) int {
	px := t.find(x)
	py := t.find(y)
	if px != py {
		return -1
	} else if t.mod[x] == t.mod[y] {
		return 1
	} else {
		return 0
	}
}

func process(n, m int) {
	var o, x, y int
	var t ufs
	t.init(n + 1)
	for i := 0; i < m; i++ {
		fmt.Scanf("%c %d %d", &o, &x, &y)
		switch o {
		case int('D'):
			t.union(x, y)
		case int('A'):
			switch t.gang(x, y) {
			case -1:
				fmt.Println("Not sure yet.")
			case 0:
				fmt.Println("In different gangs.")
			case 1:
				fmt.Println("In the same gang.")
			}
		}
	}
}

func main() {
	var t, n, m int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &m)
		process(n, m)
	}
}
