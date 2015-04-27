package sort

func QuickSort(array []int) {
	beg, end := 0, len(array)-1
	qsort(array, beg, end)
}

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

func MergeSort(array []int) {
	tmp := make([]int, len(array))
	beg, end := 0, len(array)-1
	msort(array, tmp, beg, end)
}

func msort(array, tmp []int, beg, end int) {
	if beg >= end {
		return
	}
	mid := beg + (end-beg)/2
	msort(array, tmp, beg, mid)
	msort(array, tmp, mid+1, end)
	i, j, k := beg, mid+1, beg
	for i <= mid && j <= end {
		if array[i] <= array[j] {
			tmp[k] = array[i]
			i++
			k++
		} else {
			tmp[k] = array[j]
			j++
			k++
		}
	}
	for i <= mid {
		tmp[k] = array[i]
		i++
		k++
	}
	for j <= end {
		tmp[k] = array[j]
		j++
		k++
	}
	for i := beg; i <= end; i++ {
		array[i] = tmp[i]
	}
}

func HeapSort(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		hdown(array, i, len(array))
	}
	for i := len(array) - 1; i != 0; i-- {
		xswap(array, 0, i)
		hdown(array, 0, i)
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
