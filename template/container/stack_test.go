package container_test

import (
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/container"
)

func checkStack(s *Stack, vs []interface{}) {
	assert.Must(s.Size() == len(vs))
	for i := len(vs) - 1; i >= 0; i-- {
		v, ok := s.Pop()
		assert.Must(ok && v == vs[i])
	}
	for i := 0; i < len(vs); i++ {
		s.Push(vs[i])
	}
}

func TestStack(t *testing.T) {
	var vs []interface{}
	for i := 0; i < 1024; i++ {
		vs = append(vs, i)
	}
	s := &Stack{}
	for i := 0; i < len(vs); i++ {
		s.Push(vs[i])
		checkStack(s, vs[:i+1])
	}
	for i := len(vs) - 1; i >= 0; i-- {
		v, ok := s.Pop()
		assert.Must(ok && v == vs[i])
		checkStack(s, vs[:i])
	}
}
