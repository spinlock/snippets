package permute_test

import (
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/permute"
)

func bytesToInts(bs []byte) []int {
	var is = make([]int, len(bs))
	for i := 0; i < len(bs); i++ {
		is[i] = int(bs[i])
	}
	return is
}

func intsToBytes(is []int) []byte {
	var bs = make([]byte, len(is))
	for i := 0; i < len(is); i++ {
		bs[i] = byte(is[i])
	}
	return bs
}

func test(s string, n int, f func(array []int) bool) {
	var ss = bytesToInts([]byte(s))
	mark := make(map[string]bool)
	for {
		mark[string(intsToBytes(ss))] = true
		if !f(ss) {
			break
		}
	}
	assert.Must(len(mark) == n)
	assert.Must(s == string(intsToBytes(ss)))
}

func TestPermute1(x *testing.T) {
	test("a", 1, Next)
	test("a", 1, Prev)
	test("abcdefgh", 40320, Next)
	test("hgfedcba", 40320, Prev)
}

func TestPermute2(x *testing.T) {
	test("aaa", 1, Next)
	test("aaa", 1, Prev)
	test("aaabbbcccddd", 369600, Next)
	test("dddcccbbbaaa", 369600, Prev)
}
