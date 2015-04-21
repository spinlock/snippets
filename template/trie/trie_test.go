package trie_test

import (
	"strconv"
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/trie"
)

func TestInsert(x *testing.T) {
	t := New()

	var ks []string
	var xs []string
	for i := 0; i < 32; i++ {
		ks = append(ks, strconv.Itoa(i))
		xs = append(xs, "-"+strconv.Itoa(i))
	}
	for i := 0; i < len(ks); i++ {
		t.Insert(ks[i])
		assert.Must(t.Check())
	}
	for i := 0; i < len(ks); i++ {
		assert.Must(t.Contains(ks[i]))
	}
	for i := 0; i < len(xs); i++ {
		assert.Must(t.Contains(xs[i]) == false)
	}
	assert.Must(t.Contains("") == false)
	t.Insert("")
	assert.Must(t.Check())
	assert.Must(t.Contains(""))
}

func TestInsertNoPrefix(x *testing.T) {
	t := New()

	t.Insert("000")
	assert.Must(t.InsertNoPrefix("1"))
	assert.Must(t.InsertNoPrefix("00") == false)
	assert.Must(t.InsertNoPrefix("11") == false)
	assert.Must(t.InsertNoPrefix("000") == false)
}
