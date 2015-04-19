package intree_test

import (
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/intree"
)

func makemap(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	return m
}

func minInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}

func maxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func TestInTree(t *testing.T) {
	const n = 128
	var values = make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = i //int(rand.Int31n(int32(n * n)))
	}
	var min = makemap(n)
	var max = makemap(n)
	var sum = makemap(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < i {
				min[i][j] = MaxInt
				max[i][j] = MinInt
				sum[i][j] = 0
			} else {
				min[i][j] = values[j]
				max[i][j] = values[j]
				sum[i][j] = values[j]
				if i != j {
					min[i][j] = minInt(min[i][j], min[i][j-1])
					max[i][j] = maxInt(max[i][j], max[i][j-1])
					sum[i][j] += sum[i][j-1]
				}
			}
		}
	}
	var it = New(values)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			assert.Must(min[i][j] == it.Min(i, j))
			assert.Must(max[i][j] == it.Max(i, j))
			assert.Must(sum[i][j] == it.Sum(i, j))
		}
	}
}
