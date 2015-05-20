package huffman_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/spinlock/go-libs/assert"
	. "github.com/spinlock/snippets/template/huffman"
)

func Test(x *testing.T) {
	var ss []string
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		for j := 0; j <= i; j++ {
			ss = append(ss, s)
		}
	}
	for k := len(ss) * 100; k != 0; k-- {
		i := rand.Intn(len(ss))
		j := rand.Intn(len(ss))
		ss[i], ss[j] = ss[j], ss[i]
	}
	var m = Encode(ss)
	assert.Must(len(m) == 100)
	for k := 1; k < 100; k++ {
		i := strconv.Itoa(k - 1)
		j := strconv.Itoa(k)
		assert.Must(m[i] != "")
		assert.Must(m[j] != "")
		fmt.Println(i, j, m[i], m[j])
		assert.Must(len(m[i]) <= len(m[j]))
		assert.Must(strings.HasPrefix(m[j], m[i]) == false)
	}
}
