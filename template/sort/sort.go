package sort

func QuickSort(array []int) {
	qsort(array, 0, len(array)-1)
}

func qsort(array []int, beg, end int) {
	if beg >= end {
		return
	}
	pivot := beg
	for i := beg + 1; i <= end; i++ {
		if array[i] <= array[beg] {
			pivot++
			qswap(array, i, pivot)
		}
	}
	qswap(array, beg, pivot)
	qsort(array, beg, pivot-1)
	qsort(array, pivot+1, end)
}

func qswap(array []int, i, j int) {
	if i == j {
		return
	}
	array[i], array[j] = array[j], array[i]
}

func MergeSort(array []int) {
	tmp := make([]int, len(array))
	msort(array, tmp, 0, len(array)-1)
}

func msort(array, tmp []int, beg, end int) {
	if beg >= end {
		return
	}
	mid := (beg + end) / 2
	msort(array, tmp, beg, mid)
	msort(array, tmp, mid+1, end)
	i, j := beg, mid+1
	idx := beg
	for i <= mid && j <= end {
		if array[i] <= array[j] {
			tmp[idx] = array[i]
			i++
			idx++
		} else {
			tmp[idx] = array[j]
			j++
			idx++
		}
	}
	for i <= mid {
		tmp[idx] = array[i]
		i++
		idx++
	}
	for j <= end {
		tmp[idx] = array[j]
		j++
		idx++
	}
	for i := beg; i <= end; i++ {
		array[i] = tmp[i]
	}
}
