package sort

func Search(array []int, key int) int {
	beg, end := 0, len(array)-1
	for beg <= end {
		mid := (beg + end) / 2
		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -(beg + 1)
}
