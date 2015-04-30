package main

import "fmt"

func xswap(array []int, i, j int) {
	if i != j {
		array[i], array[j] = array[j], array[i]
	}
}

func hdown(array []int, p int, size int) {
	for p < size {
		l := p*2 + 1
		r := p*2 + 2
		m := p
		if l < size && array[l] > array[m] {
			m = l
		}
		if r < size && array[r] > array[m] {
			m = r
		}
		if p == m {
			return
		}
		xswap(array, p, m)
		p = m
	}
}

func qsort(array []int, beg, end int) {
	if beg >= end {
		return
	}
	pivot := beg
	for i := beg + 1; i <= end; i++ {
		if array[i] <= array[beg] {
			pivot++
			xswap(array, i, pivot)
		}
	}
	xswap(array, beg, pivot)
	qsort(array, beg, pivot-1)
	qsort(array, pivot+1, end)
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for ; t != 0; t-- {
		var n, m int
		fmt.Scanf("%d %d", &m, &n)

		var sum = make([]int, n)
		var old = make([]int, n)
		var pls = make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &sum[i])
		}
		for i := n / 2; i >= 0; i-- {
			hdown(sum, i, n)
		}

		for l := 1; l < m; l++ {
			for i := 0; i < n; i++ {
				old[i] = sum[i]
			}
			qsort(old, 0, n-1)

			for i := 0; i < n; i++ {
				fmt.Scanf("%d", &pls[i])
			}
			qsort(pls, 0, n-1)

			for i := 0; i < n; i++ {
				sum[i] += pls[0]
			}

			for i := 1; i < n; i++ {
				for j := 0; j < n; j++ {
					x := pls[i] + old[j]
					if x >= sum[0] {
						break
					}
					sum[0] = x
					hdown(sum, 0, n)
				}
			}
		}

		qsort(sum, 0, n-1)
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", sum[i])
		}
		fmt.Println()
	}
}
