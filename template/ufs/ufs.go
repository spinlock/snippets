package ufs

func UfsInit(array []int) {
	for i := 0; i < len(array); i++ {
		array[i] = -1
	}
}

func UfsFind(array []int, x int) int {
	if px := array[x]; px < 0 {
		return x
	} else {
		array[x] = UfsFind(array, px)
		return array[x]
	}
}

func UfsSize(array []int, x int) int {
	px := UfsFind(array, x)
	return -array[px]
}

func UfsJoin(array []int, x, y int) {
	px := UfsFind(array, x)
	py := UfsFind(array, y)
	if px != py {
		array[px] += array[py]
		array[py] = px
	}
}
