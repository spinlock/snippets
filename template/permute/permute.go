package permute

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func reverse(array []int, beg, end int) {
	for i, j := beg, end; i < j; i, j = i+1, j-1 {
		swap(array, i, j)
	}
}

func NextRange(array []int, beg, end int) bool {
	var i = end - 1
	for i >= beg && array[i] >= array[i+1] {
		i--
	}
	if i >= beg {
		var j = end
		for array[i] >= array[j] {
			j--
		}
		swap(array, i, j)
	}
	reverse(array, i+1, end)
	return i >= beg
}

func Next(array []int) bool {
	return NextRange(array, 0, len(array)-1)
}

func PrevRange(array []int, beg, end int) bool {
	var i = end - 1
	for i >= beg && array[i] <= array[i+1] {
		i--
	}
	if i >= beg {
		var j = end
		for array[i] <= array[j] {
			j--
		}
		swap(array, i, j)
	}
	reverse(array, i+1, end)
	return i >= beg
}

func Prev(array []int) bool {
	return PrevRange(array, 0, len(array)-1)
}
