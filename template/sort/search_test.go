package sort_test

import (
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/sort"
)

func TestSearch(t *testing.T) {
	assert.Must(Search([]int{}, 0) == -1)

	array := make([]int, 10)
	for i := 0; i < len(array); i++ {
		array[i] = 2 * i
	}
	for i := 0; i < len(array); i++ {
		assert.Must(Search(array, i*2) == i)
		assert.Must(-(1 + Search(array, i*2-1)) == i)
		assert.Must(-(1 + Search(array, i*2+1)) == i+1)
	}
}
