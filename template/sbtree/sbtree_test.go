package sbtree_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/sbtree"
)

func checkTree(t *Tree, expect map[int]bool) {
	assert.Must(len(expect) == t.Size())
	assert.Must(t.Check())

	nodes := t.Nodes(nil)
	assert.Must(len(expect) == t.Size())
	for i := 0; i < len(nodes); i++ {
		assert.Must(t.Rank(nodes[i].Key()) == i)
		x := t.Select(i)
		assert.Must(x != nil && x.Key() == nodes[i].Key())
	}

	for i := 1; i < len(nodes); i++ {
		x := t.FindPred(nodes[i].Key())
		assert.Must(x != nil && x.Key() == nodes[i-1].Key())
	}
	for i := 0; i < len(nodes)-1; i++ {
		x := t.FindSucc(nodes[i].Key())
		assert.Must(x != nil && x.Key() == nodes[i+1].Key())
	}

	for k, _ := range expect {
		assert.Must(t.Contains(k))
	}

	for i := 1; i < len(nodes); i++ {
		assert.Must(nodes[i].Key() > nodes[i-1].Key())
	}
	for i := 0; i < len(nodes); i++ {
		assert.Must(expect[nodes[i].Key()])
	}
}

func TestInsert(x *testing.T) {
	t := New()
	checkTree(t, map[int]bool{})

	m := make(map[int]bool)
	for i := 0; i < 1024; i++ {
		m[i] = true
		_, addNode := t.Insert(i, nil)
		assert.Must(addNode)
	}
	checkTree(t, m)

	for i := 0; i < 1024; i++ {
		_, addNode := t.Insert(i, nil)
		assert.Must(addNode == false)
	}
	checkTree(t, m)
}

func TestRemove(x *testing.T) {
	t := New()
	for i := 0; i < 1024; i++ {
		_, addNode := t.Insert(i, i)
		assert.Must(addNode)
	}

	for i := 0; i < 1024; i++ {
		oldValue, delNode := t.Remove(i)
		assert.Must(delNode)
		assert.Must(oldValue.(int) == i)
	}
	checkTree(t, map[int]bool{})

	for i := 0; i < 1024; i++ {
		oldValue, delNode := t.Remove(i)
		assert.Must(delNode == false)
		assert.Must(oldValue == nil)
	}
	checkTree(t, map[int]bool{})
}

func TestRandom(x *testing.T) {
	t := New()
	m := make(map[int]bool)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 4096; i++ {
		x := int(r.Int31n(65536))
		m[x] = true
		t.Insert(x, nil)
	}
	checkTree(t, m)

	for k, _ := range m {
		delete(m, k)
		oldValue, delNode := t.Remove(k)
		assert.Must(delNode)
		assert.Must(oldValue == nil)
	}
	checkTree(t, m)
}
