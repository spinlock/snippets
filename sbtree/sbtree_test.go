package sbtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/spinlock/go-libs/assert"
)

func checkBalance(n *node) {
	if n.size == 0 {
		assert.Must(n.left.size == 0)
		assert.Must(n.right.size == 0)
	} else {
		assert.Must(n.size == n.left.size+n.right.size+1)
		assert.Must(n.left.size >= n.right.left.size)
		assert.Must(n.left.size >= n.right.right.size)
		assert.Must(n.right.size >= n.left.left.size)
		assert.Must(n.right.size >= n.left.right.size)
		checkBalance(n.left)
		checkBalance(n.right)
	}
}

func checkTree(t *SBTree, expect map[int]bool) {
	assert.Must(len(expect) == t.Size())

	if t.root != nil {
		checkBalance(t.root)
	}

	keys := t.Keys()
	assert.Must(len(expect) == t.Size())
	for i := 0; i < len(keys); i++ {
		assert.Must(t.Rank(keys[i]) == i)
		key, _, ok := t.Select(i)
		assert.Must(ok && key == keys[i])
	}

	for i := 1; i < len(keys); i++ {
		pk, _, ok := t.Predecessor(keys[i])
		assert.Must(ok && pk == keys[i-1])
	}
	for i := 0; i < len(keys)-1; i++ {
		sk, _, ok := t.Successor(keys[i])
		assert.Must(ok && sk == keys[i+1])
	}

	for k, _ := range expect {
		assert.Must(t.Contains(k))
	}

	for i := 1; i < len(keys); i++ {
		assert.Must(keys[i] > keys[i-1])
	}
	for i := 0; i < len(keys); i++ {
		assert.Must(expect[keys[i]])
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

func TestDelete(x *testing.T) {
	t := New()
	for i := 0; i < 1024; i++ {
		_, addNode := t.Insert(i, i)
		assert.Must(addNode)
	}

	for i := 0; i < 1024; i++ {
		oldValue, delNode := t.Delete(i)
		assert.Must(delNode)
		assert.Must(oldValue.(int) == i)
	}
	checkTree(t, map[int]bool{})

	for i := 0; i < 1024; i++ {
		oldValue, delNode := t.Delete(i)
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
		oldValue, delNode := t.Delete(k)
		assert.Must(delNode)
		assert.Must(oldValue == nil)
	}
	checkTree(t, m)
}
