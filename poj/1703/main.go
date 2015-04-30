package main

import "fmt"

type Ufs struct {
	size []int
	dist []int
}

func New(n int) *Ufs {
	ufs := &Ufs{}
	ufs.size = make([]int, n)
	ufs.dist = make([]int, n)
	for i := 0; i < n; i++ {
		ufs.size[i] = -1
	}
	return ufs
}

func (ufs *Ufs) Find(x int) int {
	if px := ufs.size[x]; px < 0 {
		return x
	} else {
		npx := ufs.Find(px)
		if npx != ufs.size[x] {
			ufs.size[x] = npx
			ufs.dist[x] = (ufs.dist[x] + ufs.dist[px]) % 2
		}
		return npx
	}
}

func (ufs *Ufs) Ask(x, y int) int {
	px := ufs.Find(x)
	py := ufs.Find(y)
	if px != py {
		return 0
	}
	if ufs.dist[x] != ufs.dist[y] {
		return 1
	} else {
		return 2
	}
}

func (ufs *Ufs) Diff(x, y int) {
	px := ufs.Find(x)
	py := ufs.Find(y)
	if px != py {
		ufs.size[px] += ufs.size[py]
		ufs.size[py] = px
		ufs.dist[py] = (ufs.dist[x] + 3 - ufs.dist[y]) % 2
	}
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for ; t != 0; t-- {
		var n, m int
		fmt.Scanf("%d %d\n", &n, &m)
		var ufs = New(n + 1)
		for ; m != 0; m-- {
			var op byte
			var x, y int
			fmt.Scanf("%c %d %d\n", &op, &x, &y)
			if op == 'A' {
				switch ufs.Ask(x, y) {
				case 0:
					fmt.Println("Not sure yet.")
				case 1:
					fmt.Println("In different gangs.")
				case 2:
					fmt.Println("In the same gang.")
				}
			} else if op == 'D' {
				ufs.Diff(x, y)
			}
		}
	}
}
