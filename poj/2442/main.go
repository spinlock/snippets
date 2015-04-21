package main

import "fmt"

func xswap(array []int, i, j int) {
	if i != j {
		array[i], array[j] = array[j], array[i]
	}
}

func qsort(array []int, beg, end int) {
	if beg >= end {
		return
	}
	pivot := beg
	for j := beg + 1; j <= end; j++ {
		if array[j] <= array[beg] {
			pivot++
			xswap(array, j, pivot)
		}
	}
	xswap(array, beg, pivot)
	qsort(array, beg, pivot-1)
	qsort(array, pivot+1, end)
}

func heap_down(array []int, p int) {
	size := len(array)
	for p < size {
		l := 2*p + 1
		r := 2*p + 2
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

func heap_build(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		heap_down(array, i)
	}
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for l := 0; l < t; l++ {
		var m, n int
		fmt.Scanf("%d %d", &m, &n)

		ans := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &ans[j])
		}
		qsort(ans, 0, n-1)

		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			tmp[j] = ans[j]
		}
		heap_build(tmp)

		pls := make([]int, n)
		for i := 1; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Scanf("%d", &pls[j])
			}
			qsort(pls, 0, n-1)

			for j := 0; j < n; j++ {
				tmp[j] += pls[0]
			}

			for k := 1; k < n; k++ {
				for j := 0; k < n; j++ {
					v := ans[j] + pls[k]
					if v >= tmp[0] {
						break
					}
					tmp[0] = v
					heap_down(tmp, 0)
				}
			}

			for j := 0; j < n; j++ {
				ans[j] = tmp[j]
			}
			qsort(ans, 0, n-1)
		}

		for j := 0; j < n; j++ {
			fmt.Printf("%d ", ans[j])
		}
		fmt.Println()
	}
}
