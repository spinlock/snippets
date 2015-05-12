package permute_test

import (
	"math/rand"
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

func TestPermute1(x *testing.T) {
	var ss = bytesToInts([]byte("abcdefgh"))
	for k := 0; k < len(ss)*len(ss); k++ {
		i := rand.Int31n(int32(len(ss)))
		j := rand.Int31n(int32(len(ss)))
		ss[i], ss[j] = ss[j], ss[i]
	}
	Sort(ss)
	mark := make(map[string]bool)
	for {
		mark[string(intsToBytes(ss))] = true
		if !Next(ss) {
			break
		}
	}
	assert.Must(len(mark) == 40320)
}

func TestPermute2(x *testing.T) {
	var ss = bytesToInts([]byte("abcdabcdabcd"))
	for k := 0; k < len(ss)*len(ss); k++ {
		i := rand.Int31n(int32(len(ss)))
		j := rand.Int31n(int32(len(ss)))
		ss[i], ss[j] = ss[j], ss[i]
	}
	Sort(ss)
	mark := make(map[string]bool)
	for {
		mark[string(intsToBytes(ss))] = true
		if !Next(ss) {
			break
		}
	}
	assert.Must(len(mark) == 369600)
}
