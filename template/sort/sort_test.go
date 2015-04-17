package sort_test

import (
	"math/rand"
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/sort"
)

func newArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	for x := n * n; x != 0; x-- {
		i := int(rand.Int31n(int32(n)))
		j := int(rand.Int31n(int32(n)))
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func checkSort(fsort func([]int)) {
	for n := 0; n <= 10; n++ {
		array := newArray(1 << uint(n))
		fsort(array)
		for i := 0; i < len(array); i++ {
			assert.Must(array[i] == i)
		}
	}
}

func TestQuickSort(t *testing.T) {
	checkSort(QuickSort)
}

func TestMergeSort(t *testing.T) {
	checkSort(MergeSort)
}

func TestHeapSort(t *testing.T) {
	checkSort(HeapSort)
}
