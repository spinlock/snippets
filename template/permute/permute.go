package permute

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
	xswap(array, pivot, beg)
	qsort(array, beg, pivot-1)
	qsort(array, pivot+1, end)
}

func xswap(array []int, i, j int) {
	if i != j {
		array[i], array[j] = array[j], array[i]
	}
}

func Sort(array []int) {
	qsort(array, 0, len(array)-1)
}

func Next(array []int) bool {
	var part = len(array) - 2
	for part >= 0 && array[part] >= array[part+1] {
		part--
	}
	if part < 0 {
		return false
	}
	var xchg = len(array) - 1
	for array[xchg] <= array[part] {
		xchg--
	}
	xswap(array, xchg, part)
	for i, j := part+1, len(array)-1; i < j; i, j = i+1, j-1 {
		xswap(array, i, j)
	}
	return true
}
