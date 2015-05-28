package main

import (
	"bufio"
	"fmt"
	"os"
)

func split(b []byte) [][]byte {
	var ss [][]byte
	var n = len(b)
	for l, r := 0, 0; r <= n; r++ {
		if r == n || b[r] == ' ' {
			if l < r {
				ss = append(ss, b[l:r])
			}
			l = r + 1
		}
	}
	return ss
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func process(input []byte) int {
	var ss = split(input)
	if len(ss) != 2 {
		return 0
	}
	var a, b = ss[0], ss[1]
	var m, n = len(a), len(b)
	var so = make([]int, n+1)
	var sn = make([]int, n+1)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				sn[j] = 1 + so[j-1]
			} else {
				sn[j] = max(so[j], sn[j-1])
			}
		}
		so, sn = sn, so
	}
	return so[n]
}

func main() {
	var r = bufio.NewReader(os.Stdin)
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			return
		}
		fmt.Println(process(b))
	}
}
