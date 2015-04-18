package heap_test

import (
	"math/rand"
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/heap"
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

func TestHeap(t *testing.T) {
	for i := 0; i <= 10; i++ {
		array := newArray(1 << uint(i))
		h := New()
		for j := 0; j < len(array); j++ {
			h.Push(array[j])
		}
		for j := 0; j < len(array); j++ {
			v, ok := h.Pop()
			assert.Must(ok && v == j)
		}
		assert.Must(h.Size() == 0)
	}
}
