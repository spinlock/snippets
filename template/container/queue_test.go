package container_test

import (
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/container"
)

func checkQueue(q *Queue, vs []interface{}) {
	assert.Must(q.Size() == len(vs))
	for i := 0; i < len(vs); i++ {
		v, ok := q.Pop()
		assert.Must(ok && v == vs[i])
	}
	for i := 0; i < len(vs); i++ {
		q.Push(vs[i])
	}
}

func TestQueue(t *testing.T) {
	var vs []interface{}
	for i := 0; i < 1024; i++ {
		vs = append(vs, i)
	}
	q := &Queue{}
	for i := 0; i < len(vs); i++ {
		q.Push(vs[i])
		checkQueue(q, vs[:i+1])
	}
	for i := 0; i < len(vs); i++ {
		v, ok := q.Pop()
		assert.Must(ok && v == vs[i])
		checkQueue(q, vs[i+1:])
	}
}
